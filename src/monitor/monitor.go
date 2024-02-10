package monitor

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	_ "github.com/lib/pq"
)

var (
	metricName      = ""
	aggregationType = ""
	resourceUri     = ""
)

func insertMonitorData(db *sql.DB, data string) error {
	sqlStatement := `
        INSERT INTO monitor_data (data, timestamp)
        VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, data, time.Now())
	return err
}

func fetchMetrics() (string, error) {
	url := "https://management.azure.com/" + resourceUri + "/providers/microsoft.insights/metrics?api-version=2018-01-01&timespan=P1D&interval=P1D&metricnames=" + metricName + "&aggregation=" + aggregationType + ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to create credential: %v", err)
	}

	tk, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{"https://management.azure.com/.default"}})
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	fmt.Printf("Access token: %s\n", tk.Token)

	req.Header.Add("Authorization", "Bearer "+tk.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func generateSasToken(resourceUri string, keyName string, key string) (string, error) {
	// Set token valid period
	sinceEpoch := time.Now().Unix()
	expiry := sinceEpoch + 60*60*24*7 // 1 week

	// String to sign
	stringToSign := fmt.Sprintf("%s\n%d", url.QueryEscape(resourceUri), expiry)

	// Generate the HMAC
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(stringToSign))

	// Create the SAS token
	sasToken := fmt.Sprintf("SharedAccessSignature sr=%s&sig=%s&se=%d&skn=%s",
		url.QueryEscape(resourceUri),
		url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil))),
		expiry,
		url.QueryEscape(keyName),
	)

	return sasToken, nil
}

func fetchInstallationInfo(installationId string) (string, error) {
	url := fmt.Sprintf("https://%s.servicebus.windows.net/%s/installations/%s?api-version=2017-04", os.Getenv("HUB_NAMESPACE"), os.Getenv("HUB_NAME"), installationId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resourceUri := os.Getenv("RESOURCE_URI")
	keyName := os.Getenv("KEY_NAME")
	key := os.Getenv("KEY")

	sasToken, err := generateSasToken(resourceUri, keyName, key)
	if err != nil {
		fmt.Println("Error generating SAS token:", err)
	}

	req.Header.Add("Authorization", "SharedAccessSignature "+sasToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	http.HandleFunc("/api/InsertMonitorData", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))
		if err != nil {
			http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Fetch data from Azure Monitor API...
		data, err := fetchMetrics()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = insertMonitorData(db, data)
		if err != nil {
			http.Error(w, "Failed to insert monitor data", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Monitor data inserted successfully")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
