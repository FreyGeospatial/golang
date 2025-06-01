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
	default:
		result = "It's the weekend!"
	}
	fmt.Println(result)

	// can also declare variable in switch statement (can also do this in if statement).
	// the advantage is that the variable is scoped to the switch statement only.
	switch day := weekday.String(); day {
	case "Sunday":
		fmt.Println("It's Sunday, time to relax!")
	case "Monday":
		fmt.Println("It's Monday, back to work!")
	case "Tuesday":
		fmt.Println("It's Tuesday, keep going!")
	case "Wednesday":
		fmt.Println("It's Wednesday, we're halfway there!")
	case "Thursday":
		fmt.Println("It's Thursday, almost the weekend!")
	case "Friday":
		fmt.Println("It's Friday, the weekend is near!")
	case "Saturday":
		fmt.Println("It's Saturday, enjoy your day off!")
	default:
		fmt.Println("Unknown day, please check your calendar!")
	}
	// fmt.Println(day) // Uncommenting this line will cause an error because 'day' is scoped to the switch statement only.

	// normally, switch statements will break after a case is matched,
	// but you can use fallthrough to continue to the next case.
	switch num := 6; {
	case num%2 == 0:
		fmt.Println(num, "is an even number")
		fallthrough
	case (num % 2) != 0:
		fmt.Println(num, "is an odd number, but using fallthrough here will print this regardless of the previous case.")
	default:
		fmt.Println(num, "is zero or not a valid number. If one of the previous cases matched, this will not be printed, unless you use fallthrough on the previous case.")
	}
}
