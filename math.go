package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum) //i float issue with precision

	total := float64(intSum) + floatSum
	fmt.Println("Total sum:", total)

	round := math.Round(floatSum*100) / 100
	fmt.Println("Rounded float sum:", round)

	fmt.Println(("The value of pi is: "), math.Pi)

}
