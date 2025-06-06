package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, 10, 23, 0, 0, 0, 0, time.UTC) // date golang was formally announced
	fmt.Println("Go was announced on: %v", t)

	fmt.Println("Current time is: ", time.Now())
	fmt.Printf("This object is of type: %T\n", t)

	// create time object by parsing a string
	now := time.Now()
	now_string := now.Format(time.ANSIC)
	fmt.Println("Current time in ANSIC format: ", now_string)

	tomorrow := now.AddDate(0, 0, 1) // add one day to the current date
	fmt.Println("Tomorrow's date is: ", tomorrow.Format(time.ANSIC))

	format := "2006-01-02"              // must use a specific date format to format the date, unlike in other languages
	fmt.Printf(tomorrow.Format(format)) // format the date in a custom format
}
