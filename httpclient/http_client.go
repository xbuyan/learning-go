package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	// Step 1: Send GET request

	resp, err := http.Get("https://httpbin.org/get")

	if err != nil {

		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	//Step 2: Check status code

	fmt.Println("Status:", resp.StatusCode)

	//Step 3: Read the body

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		fmt.Println("Error reading response:", err)
		os.Exit(1)

	}

	fmt.Println(string(body))
	fmt.Println("\n--- POST Request ---")
	postRecord()
}

func postRecord() {

	// Step 1. Create a sample FileRecord

	record := map[string]interface{}{

		"filename":    "evidence.jpg",
		"size":        204800,
		"sha256":      "7548bd8ac2804a976da68762428fdbb6190f7a23e5d537ff7d5386ad184cfe06",
		"recorded_at": "2026-04-18 10:00:00",
	}

	// Step 2: Convert to JSON bytes

	jsonBytes, err := json.Marshal(record)

	if err != nil {

		fmt.Println("Error marshalling", err)
		os.Exit(1)
	}

	//Step 3: Send post request

	resp, err := http.Post(
		"https://httpbin.org/post",
		"application/json",
		bytes.NewReader(jsonBytes),
	)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Step 4: Read and print response

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response", err)
		os.Exit(1)
	}
	fmt.Println("Status:", resp.StatusCode)
	fmt.Println(string(body))
}
