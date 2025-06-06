// go is organized with packages, and packages are organized into functions.
// This is the main package, which is the entry point of the Go program.

// furthermore, the convention in go is to use camelCase for function names.
// Case is important in Go, and the first letter of a function name determines its visibility.
// If the first letter is uppercase, the function is exported and can be accessed from other packages.
// If the first letter is lowercase, the function is unexported and can only be accessed within the same package.
package main

import (
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something...")
	count, sum, average := summaryStats(5, 10, 15, 20, 25)
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Average: %.2f\n", average)
}

func addValues(value1 int, value2 int) int {
	// This function adds two integers and returns the result.
	return value1 + value2
}

func addValuesWithOneType(value1, value2 int) int {
	// This function adds two integers and returns the result.
	// It uses a single type for both parameters.
	return value1 + value2
}

func addArbitraryValues(values ...int) int {
	// This function adds an arbitrary number of integers and returns the result.
	// The ... syntaxc is read in as a slice of integers.
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// in go, we can also return multiple values from a function.
func summaryStats(values ...int) (int, int, float64) {
	// This function returns the count, sum and average of an arbitrary number of integers.
	sum := 0
	for _, value := range values {
		sum += value
	}
	count := len(values)
	average := float64(sum) / float64(count)
	return count, sum, average
}
