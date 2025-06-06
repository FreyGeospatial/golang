package main

// Uncomment this import section to use required Go packages
import (
	"bufio"
	"fmt"
	"os"
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
		panic("Value1 is not a number")
	}
	num2, err2 := strconv.ParseFloat(value2, 64) // convert the string to a float
	if err2 != nil {
		fmt.Println("Error converting input to float:", err2)
		panic("Value2 is not a number")
	}

	return num1 + num2
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Good practice to use bufio.NewReader for reading input
	fmt.Print("Enter a number: ")
	val1, _ := reader.ReadString('\n') // generally better than fmt.Scan or fmt.Scanln, which can be tricky with spaces

	fmt.Print("Enter another number: ")
	val2, _ := reader.ReadString('\n')

	result := calculate(val1, val2)             // call the calculate function
	fmt.Printf("The result is: %.2f\n", result) // print the result
}
