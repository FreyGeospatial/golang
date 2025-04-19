package main

import (
	"fmt"
	"strconv"
)

func main() {

	var name string
	var age_string string
	var hourly_wage_string string
	var hours_string string

	fmt.Print("Enter your name:")
	fmt.Scan(&name)
	fmt.Print("Enter your age:")
	fmt.Scan(&age_string)
	fmt.Print("Enter your hourly wage:")
	fmt.Scan(&hourly_wage_string)
	fmt.Print("Enter the number of hours you work per week, on average:")
	fmt.Scan(&hours_string)

	fmt.Println()
	fmt.Println("Your name is ", name)

	age_int, err := strconv.Atoi(age_string)

	// parsing conversions will always return an error variable
	// if the conversion fails, so we need to check for that
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	hourly_wage_float, err := strconv.ParseFloat(hourly_wage_string, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}

	hours_float, err := strconv.ParseFloat(hours_string, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}

	weekly_earnings := hourly_wage_float * hours_float

	// %d for int, %f for float, %s for string
	fmt.Println(fmt.Sprintf("Your age is %d", age_int)) // not necessary of course to do this conversion, but just to show
	fmt.Println(fmt.Sprintf("Your hourly wage is %f", hourly_wage_float))
	fmt.Println(fmt.Sprintf("The number of hours you work per week is %f", hours_float))
	fmt.Println(fmt.Sprintf("You earn $%.2f per week", weekly_earnings)) // 2 decimal places
}
