package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Advanced Calculator!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first value: ")
	val1, _ := reader.ReadString('\n')
	val1_float := convertStringToFloat(val1)
	fmt.Print("Enter the operation (+, -, *, /): ")
	operation, _ := reader.ReadString('\n')
	fmt.Print("Enter second value: ")
	val2, _ := reader.ReadString('\n')
	val2_float := convertStringToFloat(val2)
	result := calc(val1_float, val2_float, strings.TrimSpace(operation))
	fmt.Printf("The result of %.2f %s %.2f is: %.2f\n", val1_float, strings.TrimSpace(operation), val2_float, result)
	fmt.Println("Thank you for using the Advanced Calculator!")

}

func calc(val1 float64, val2 float64, operation string) float64 {
	if operation == "+" {
		return addValues(val1, val2)
	}
	if operation == "-" {
		return subtractValues(val1, val2)
	}
	if operation == "*" {
		return multiplyValues(val1, val2)
	}
	if operation == "/" {
		return divideValues(val1, val2)
	}
	panic("Error: Invalid operation. Please use +, -, *, or /.")
}

func convertStringToFloat(value string) float64 {
	// Convert string to float64
	value = strings.TrimSpace(value)
	my_float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("Error converting string to float")
	}
	return my_float
}

func addValues(val1 float64, val2 float64) float64 {
	return val1 + val2
}

func subtractValues(val1 float64, val2 float64) float64 {
	return val1 - val2
}

func multiplyValues(val1 float64, val2 float64) float64 {
	return val1 * val2
}

func divideValues(val1 float64, val2 float64) float64 {
	if val2 == 0 {
		panic("Error: Division by zero is not allowed.")
	}
	return val1 / val2
}
