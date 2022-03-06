package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func return_json(w http.ResponseWriter, req *http.Request) {
	// Dummy JSON object
	var payload = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	// Set header for response
	w.Header().Set("Content-Type", "application/json")
	// Encode JSON object into response
	json.NewEncoder(w).Encode(payload)
}

func main() {
	var port = 8090

	// Initialize handlers
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/return_json", return_json)

	// Start http server
	log.Println("Starting server at port:", port)
	http.ListenAndServe(":8090", nil)
}
