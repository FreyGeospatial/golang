package main

import "fmt"

var first_name string
var last_name string

func main() {
	fmt.Print("Enter your first name: ")
	fmt.Scan(&first_name)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&last_name)

	fmt.Println()
	fmt.Println("Hello, " + first_name + " " + last_name + "!")
	fmt.Println()
}
