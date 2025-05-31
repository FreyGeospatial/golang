package main

import (
	"fmt"
	"sort"
)

// Slices in Go are more flexible than arrays and can grow or shrink in size, and be sorted.
// They are built on top of arrays and provide a dynamic view of the underlying array.

func main() {
	names_array := [3]string{"Jordan", "James", "Joe"} // Declare an array of strings with a fixed size of 3 and Assign values to the array
	fmt.Println("Names array:", names_array)           // Print the entire array

	var colors_slice = []string{"Red", "Green", "Blue", "Orange"} // Declare and initialize a slice of strings
	fmt.Println("Initial slice:", colors_slice)                   // Print the initial slice

	// a more memory-efficient way to create a slice is to use the make function, which declares a slice and allocates memory for it.
	colors_slice2 := make([]string, 0, 5)                  // Create a slice with the initial length of 0 and a capacity of 5
	fmt.Println("Slice created with make:", colors_slice2) // Print the slice created with make
	// add elements to the slice using the append function, which adds elements to the end of the slice and returns a new slice with the added elements.
	colors_slice2 = append(colors_slice2, "Purple", "Yellow", "Magenta", "Orange", "Blue") // Append new colors to the slice
	fmt.Println("Slice after appending:", colors_slice2)

	// I can add more elements than the initial capacity, and Go will automatically resize the underlying array if needed.
	colors_slice2 = append(colors_slice2, "Cyan", "Pink")         // Append more colors to the slice
	fmt.Println("Slice after adding more colors:", colors_slice2) // Print the slice after adding more colors

	colors_slice2 = remove_element(colors_slice2, 3)                       // Remove the element at index 3 (Orange)
	fmt.Println("Slice after removing element at index 3:", colors_slice2) // Print the slice after removing an element

	sort.Strings(colors_slice2)                 // Sort the slice in alphabetical order
	fmt.Println("Sorted slice:", colors_slice2) // Print the sorted slice
}

func remove_element(slice []string, index int) []string {
	// return a slice up to the specified index and concatenate it with the slice from the next index onward.
	return append(slice[:index], slice[index+1:]...) // Remove the element at the specified index
}
