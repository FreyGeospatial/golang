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
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n') // generally better than fmt.Scan or fmt.Scanln, which can be tricky with spaces
	fmt.Println("You entered:", str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')                           // do not need to re-declare str as it is already declared
	num, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // parse the string to a float. Need to trim the newline character
	if err != nil {
		fmt.Println("Error converting input:", err)
	} else {
		fmt.Printf("You entered: %.2f\n", num)
	}
}
