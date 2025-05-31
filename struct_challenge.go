// Code Challenge: Convert slices of simple values to a slice of objects

package main

// youre given a slice of strings representing colors and a slice of integers representing their hex representation.
// Youre also given a definition of struct named Color.

// Return a slice containing instances of Color, where each instance is created from the corresponding elements in the colors and hex values slices.

func slices_to_objects(color_names []string, hexValues []int) []Color {
	// Create a slice of Color objects
	my_slice := make([]Color, 0)
	for i := 0; i < len(color_names); i++ {
		// Create a new Color object for each color name and hex value
		color := Color{
			Name: color_names[i],
			Hex:  hexValues[i],
		}
		// Append the Color object to the slice
		my_slice = append(my_slice, color)
	}
	return my_slice // Return the slice of Color objects
}

type Color struct {
	Name string
	Hex  int
}

func main() {
	color_names := []string{"Red", "Green", "Blue"}
	hex_values := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slices_to_objects(color_names, hex_values)
	for _, color := range colors {
		println("Color Name:", color.Name, "Hex Value:", color.Hex)
	}
}
