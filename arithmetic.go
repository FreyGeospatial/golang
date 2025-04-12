package main

import "fmt"

// should just use float64 most of the time, as it is more precise
var rate_of_pay float64 = 17.25
var hours_worked float64 = 40

var gross_pay = rate_of_pay * hours_worked

func main() {
	fmt.Println(fmt.Sprintf("Rate of pay is: %f", rate_of_pay))
	fmt.Println(fmt.Sprintf("Hours worked is: %f", hours_worked))
	fmt.Println(fmt.Sprintf("Gross pay is: %f", gross_pay))
}
