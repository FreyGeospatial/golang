package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	fmt.Println("Today is", weekday)
	day_number := int(weekday)
	fmt.Println("Day number:", day_number)

	var result string
	switch day_number {
	case 0:
		result = "It's a Sunday!"
	case 1:
		result = "It's a Monday!"
	case 2:
		result = "It's a Tuesday!"
	case 3:
		result = "It's a Wednesday!"
	case 4:
		result = "It's a Thursday!"
	case 5:
		result = "It's a Friday!"
	case 6:
		result = "It's a Saturday!"
	default:
		result = "Invalid day number!"
	}
	fmt.Println(result)
}
