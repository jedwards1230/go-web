package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func client() {
	// Parse cli arguments
	ip := flag.String("ip", "127.0.0.1", "IP for the client to connect to")
	port := flag.Int("port", 8090, "Port for the client to connect to")
	dir := flag.String("dir", "", "Director to connect to")
	flag.Parse()

	// Get response from host
	fmt.Println(fmt.Sprintf("Connecting to %v:%d%v", *ip, *port, *dir))
	fmt.Println("Sending request...")
	resp, err := http.Get(fmt.Sprintf("http://%v:%d/%v", *ip, *port, *dir))
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
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(scanner.Text()), &jsonMap)
		fmt.Println("Key 1:", jsonMap["key1"])
		fmt.Println("Key 2:", jsonMap["key2"])
	}

	// Panic if the scanner throws an error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
