package main

import (
	"fmt"
)

func main() {
	// Declare and initialize variables
	var (
		name    string  = "John Doe"
		age     int     = 30
		height  float64 = 5.9
		isAdult bool    = true
	)

	// Interpolate variables into a string and print at same time
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Adult: %t\n", name, age, height, isAdult)
	// can also use %v for default format
	fmt.Printf("Name: %v, Age: %v, Height: %.1f, Is Adult: %v\n", name, age, height, isAdult)

	// Using fmt.Sprintf to create a formatted string
	formattedString := fmt.Sprintf("Name: %s, Age: %d, Height: %.1f, Is Adult: %t", name, age, height, isAdult)
	fmt.Println(formattedString)
}
