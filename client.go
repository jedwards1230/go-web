package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get response from host
	fmt.Println("Sending request...")
	resp, err := http.Get("http://localhost:8090/return_json")
	// Panic if error
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create scanner object to read http response
	fmt.Println("Response Status: ", resp.Status)
	scanner := bufio.NewScanner(resp.Body)

	// Print each line the scanner finds
	for i := 0; scanner.Scan() && i < 5; i++ {
		var obj = scanner.Text()
		fmt.Println(obj)
	}

	// Panic if the scanner throws an error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
