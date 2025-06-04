package main

import (
	"fmt"
)

// youre given a map of string keys and associated instances of a struct called Item:

func main() {
	apple := Item{Name: "Apple", Price: 0.5, Quantity: 3}
	orange := Item{Name: "Orange", Price: 0.75, Quantity: 2}
	carrot := Item{Name: "Carrot", Price: 0.25, Quantity: 5}
	cart = append(cart, apple, orange, carrot)              // Add items to the cart
	total := calculate_total(cart)                          // Calculate the total value of the cart
	fmt.Println("Total value of the shopping cart:", total) // Print the total value of the shopping cart
}

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

var cart = []Item{}

func calculate_total(cart []Item) float64 {
	total := 0.0
	for _, item := range cart {
		total = total + (float64(item.Quantity) * item.Price)
	}
	return total
}
