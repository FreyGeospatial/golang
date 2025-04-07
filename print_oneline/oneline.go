package main

import "fmt"

// In Go, := is for declaration + assignment, whereas = is for assignment only.

// For example, var foo int = 10 is the same as foo := 10.

func main() {
	message := fmt.Sprintf("I am a %s-liner!", "one")
	fmt.Println(message)
}
