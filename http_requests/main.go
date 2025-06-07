package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	content := string(bytes)
	tours := toursFromJSON(content)
	for _, tour := range tours {
		fmt.Printf("Tour: %s, Price: %s\n", tour.Name, tour.Price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Tour struct {
	Name, Price string
}

func toursFromJSON(content string) []Tour {
	// This function would parse the JSON data and return a slice of Tour structs.
	// For simplicity, we will return an empty slice here.
	tours := make([]Tour, 0)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token() // Read the opening bracket
	checkError(err)
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour) // Decode each tour into the Tour struct
		checkError(err)
		tours = append(tours, tour) // Append the tour to the slice
	}
	return tours
}
