package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the total seconds: ")
	total_seconds_string, _ := reader.ReadString('\n')

	// convert the string to a float
	total_seconds, err := strconv.ParseFloat(strings.TrimSpace(total_seconds_string), 64)
	if err != nil {
		fmt.Println("Error converting input to float:", err)
		return
	}
	total_hours := total_seconds / 3600.0
	minutes := int(total_seconds) / 60
	remaining_seconds := int(total_seconds) % 60
	fmt.Printf(("Total hours: %.2f\n"), total_hours)
	fmt.Printf(("Or... %d minutes, %d seconds\n"), minutes, remaining_seconds)
}
