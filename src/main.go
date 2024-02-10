package main

import (
	"azurescaler/v3/create"
	"azurescaler/v3/delete"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type VMSet struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func postHandler(w http.ResponseWriter, r *http.Request) []byte {
	var vmSet VMSet
	err := json.NewDecoder(r.Body).Decode(&vmSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	delete.DeleteAzure(vmSet.Name)

	resByte, _ := json.Marshal("Deletion was successful")
	return resByte
}

func main() {
	listenAddr := ":8080"

	http.HandleFunc("/api/HttpTriggerAzureScaler", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go! -- We will start the magic sooon")
		create.CreateAzure("start") // Ensure this function exists in the "create" package
		fmt.Fprint(w, "Hello from Go! --msgic show completed")
		var resByte []byte
		resByte, _ = json.Marshal("Creation and Deletion was successful")
		w.Header().Set("Content-Type", "application/json")
		w.Write(resByte)

	})

	http.HandleFunc("/api/HttpTriggerAzureScaleDown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go! -- We will start the magic sooon")
		var resByte []byte

		switch r.Method {
		case http.MethodPost:
			resByte = postHandler(w, r)

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resByte)
		fmt.Fprint(w, "Hello from Go! --msgic show completed")

	})

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))

}
