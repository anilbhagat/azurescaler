package main

import (
	"azurescaler/v3/create"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go!")
		create.CreateAzure("start") // Ensure this function exists in the "create" package

	})

	http.ListenAndServe(":8080", nil)

}
