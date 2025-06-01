package main

import (
	"fmt"
)

func main() {
	the_answer := 42
	if the_answer == 42 {
		fmt.Println("The answer is 42!")
	} else if the_answer < 42 {
		fmt.Println("The answer is less than 42.")
	} else {
		fmt.Println("The answer is greater than 42.")
	}
}
