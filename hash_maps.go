package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps in Go are unordered collections of key-value pairs.")

	// instantiate a map using the make function
	states := make(map[string]string)   // first argument is the key type, second is the value type
	fmt.Println("Initial map:", states) // Print the initial empty map
	states["CA"] = "California"         // Add key-value pairs to the map
	states["TX"] = "Texas"
	states["NY"] = "New York"
	states["NJ"] = "New Jersey"
	states["FL"] = "Florida"

	fmt.Println("Map after adding some states:", states) // Print the map after adding some states

	cali := states["CA"]                     // Access a value by its key
	fmt.Println("Value for key 'CA':", cali) // Print the value for the key "CA"

	// remove a key-value pair from the map
	delete(states, "TX")                                // Remove the key "TX" from the map
	fmt.Println("Map after deleting key 'TX':", states) // Print the map after deletion

	// loop through the map using a for range loop
	for key, value := range states {
		fmt.Printf("Key: %s, Value: %s\n", key, value) // Print each key-value pair in the map
	}

	// create slice of keys from the map
	keys := make([]string, len(states)) // Create a slice to hold the keys. you don't need to specify the size, but it's more efficient to preallocate it with the length of the map.
	i := 0
	for key := range states {
		keys[i] = states[key]
		i++
	}
	sort.Strings(keys)
	fmt.Println("Keys in the map:", keys) // Print the slice of keys
}
