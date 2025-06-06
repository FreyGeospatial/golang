package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Good practice to use bufio.NewReader for reading input
	fmt.Print("Enter a value: ")
	str1, _ := reader.ReadString('\n')
	fmt.Print("Enter another value: ")
	str2, _ := reader.ReadString('\n')
	val1, err1 := strconv.ParseFloat(strings.TrimSpace(str1), 64)
	val2, err2 := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting one of the input:", err1, err2)
	}
	result := val1 + val2
	fmt.Printf("The result is: %.2f\n", result)
}
