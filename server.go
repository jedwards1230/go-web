package main

import (
	"encoding/json"
	"flag"
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
	// Parse cli arguments
	ip := flag.String("ip", "127.0.0.1", "IP for the client to connect to")
	port := flag.Int("port", 8090, "Port for the server to listen on")
	flag.Parse()

	// Initialize handlers
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/return_json", return_json)

	// Start http server
	log.Println(fmt.Sprintf("Starting server at %v:%d", *ip, *port))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
