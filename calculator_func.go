package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	value1 = strings.TrimSpace(value1)           // remove any leading or trailing whitespace
	value2 = strings.TrimSpace(value2)           // remove any leading or trailing whitespace
	num1, err1 := strconv.ParseFloat(value1, 64) // convert the string to a float
	if err1 != nil {
		fmt.Println("Error converting input to float:", err1)
	}
	num2, err2 := strconv.ParseFloat(value2, 64) // convert the string to a float
	if err2 != nil {
		fmt.Println("Error converting input to float:", err2)
	}

	return num1 + num2
}
