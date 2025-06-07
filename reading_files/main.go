package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("This program demonstrates how to read from a file in Go.")
	my_filepath := "./fromString.txt" // Define the file path
	file, err := os.Create(my_filepath)
	defer file.Close() // wait until the function returns to close the file
	checkError(err)
	fmt.Println("File content:")
	fmt.Println()
	length, err := io.WriteString(file, "Hello, from Go!\n")
	fmt.Println()
	fmt.Printf("Write a file with %v characters\n", length)

	readFile(my_filepath) // Call the readFile function to read the file
	fmt.Println("File reading completed successfully.")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close() // wait until the function returns to close the file
	checkError(err)    // Check for errors when opening the file

	content, err := io.ReadAll(file)
	checkError(err) // Check for errors when reading the file

	fmt.Println(string(content))
}
