// read shopping cart from a JSON file
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	if err != nil {
		panic("Error reading JSON content: " + err.Error())
	}
	for decoder.More() {
		var item CartItem
		err := decoder.Decode(&item) // Decode each item into the Item struct
		if err != nil {
			panic("Error decoding JSON item: " + err.Error())
		}
		cart = append(cart, item) // Append the item to the cart
	}
	return cart
}

func main() {
	fmt.Println("This program reads a shopping cart from a JSON file.")
	content := `[{"name":"apple","price":2.99,"quantity":3}, {"name":"orange","price":1.50,"quantity":8}, {"name":"banana","price":0.49,"quantity":12}]`
	cart := getCartFromJson(content)

	fmt.Printf("Updated Shopping Cart: %+v\n", cart) // Print the updated shopping cart
}
