package main

import (
	"fmt"
)

func main() {
	string_length, err := fmt.Println("Hello, World!") // note that you must use the variables you declare... otherwise the compiler will complain
	if err == nil {
		fmt.Println("String length:", string_length) // note that this will count an extra character for the newline
	}
}
