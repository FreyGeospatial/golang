package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Sound string
}

func main() {
	dog := Dog{
		Breed: "Labrador",
		Sound: "Woof",
	}

	fmt.Println("The dog is a", dog.Breed, "and it says", dog.Bark())
}

func (d Dog) Bark() string {
	return d.Sound
}
