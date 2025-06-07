package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("This program demonstrates how to make HTTP requests in Go.")
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "") // Set a user agent to avoid being blocked by some servers
	checkError(err)                      // Check for errors when creating the request
	response, err := client.Do(request)
	checkError(err)             // Check for errors when making the request
	defer response.Body.Close() // Ensure the response body is closed after we're done
	fmt.Printf("Response type: %s\n", response.Header.Get("Content-Type"))
	fmt.Printf("Response status: %s\n", response.Status)

	bytes, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Response body:\n\n")
	fmt.Println(string(bytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
