package main

import (
	"fmt"
)

func main() {
	fmt.Println(("Hello, World!"))
	colors := []string{"Red", "Green", "Blue"}

	// one time of loop
	for i, color := range colors {
		fmt.Printf("Color %d: %s\n", i+1, color)
	}
	// another way to loop through a slice
	for i := range colors {
		fmt.Printf("Color %d: %s\n", i+1, colors[i])
	}

	// yet another way to loop through a slice
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Color %d: %s\n", i+1, colors[i])
	}

	states := make(map[string]string)
	states["CA"] = "California"
	states["TX"] = "Texas"
	for state, name := range states {
		fmt.Printf("State: %s, Name: %s\n", state, name)
	}

	value := 0
	sum := 0
	for value < 10 { // equivalent to while loop, which does not have a dedicated syntax in Go
		sum += value
		value++
		fmt.Println("Current value:", value, "Sum so far:", sum)
	}

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
		if sum2 > 200 {
			goto the_end
		}
	}
	fmt.Println("Final sum:", sum2) // this line will not be executed if the goto statement is reached
the_end:
	println("Goto allows you to jump to an arbitrary label anywhere in the code, but it's generally discouraged in favor of structured control flow.")

	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum2
		if sum3 > 200 {
			break
		}
	}
	fmt.Println("Final sum3:", sum3) // this line will be executed if the loop is broken
}
