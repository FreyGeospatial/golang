package main

import "fmt"

// Arrays in Go are fixed-size sequences of elements of the same type.
// You can't add or remove elements from an array after it is created, and its size is determined at compile time.

// for more flexibility, you can use slices, which are built on top of arrays and can grow or shrink in size.
func main() {
	var colors [3]string // Declare an array of strings with a fixed size of 3
	colors[0] = "Red"    // Assign values to the array
	colors[1] = "Green"
	colors[2] = "Blue"                     // Assign values to the array
	fmt.Println("Colors array:", colors)   // Print the entire array
	fmt.Println("First color:", colors[0]) // Access and print the first element
	// iterate over the array using a for loop
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Color at index %d: %s\n", i, colors[i]) // Print each color with its index
	}

	array2 := [5]int{1, 2, 3, 4, 5}           // Declare and initialize an array of integers
	fmt.Println("Array of integers:", array2) // Print the entire array of integers

	fmt.Println("Length of colors array:", len(colors)) // Print the length of the colors array
}
