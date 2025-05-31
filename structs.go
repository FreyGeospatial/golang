package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in Go are composite data types that group together variables (fields) under a single name.")
	my_dog := Dog{ // Create an instance of the Dog struct
		Name:   "Rudy",  // Assign values to the fields
		Age:    5,       // Age of the dog
		Breed:  "Boxer", // Breed of the dog
		Weight: 30.5,    // Weight of the dog in kilograms
	}

	fmt.Println("my_dog struct:", my_dog)   // Access and print the Age field
	fmt.Println("Dog's Name:", my_dog.Name) // Access and print the Name field
}

type Dog struct { // capitalized name means it's exported and can be accessed from other packages
	// Dog represents a dog with its attributes. it is sort of like a class in other languages.
	// However, Go does not have classes (there is no concept of inheritcance) but structs can be used to achieve similar functionality.
	Name   string  // Name of the dog
	Age    int     // Age of the dog in years
	Breed  string  // Breed of the dog
	Weight float64 // Weight of the dog in kilograms
}
