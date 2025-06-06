package main

import (
	"fmt"
	"time"
)

func main() {
	// turn function call into a goroutine
	go say("Hello, Goroutine!")
	fmt.Println("Hello, from main!")

	// anonymous function/goroutine
	go func(message string) {
		time.Sleep(1 * time.Second) // simulate some work
		fmt.Println(message)
	}("Hello, anonymous Goroutine!")

	time.Sleep(2 * time.Second) // wait for goroutine to finish
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
