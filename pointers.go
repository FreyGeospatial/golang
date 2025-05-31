package main

func is_valid_pointer(value *int) bool {
	if value == nil {
		return false
	} else {
		return true
	}

}

func main() {
	// Example of using pointers in Go.
	//  a pointer is a variable that stores the memory address of another value rather than the value itself.

	// Pointers matter because they allow you to:
	// 1. Share data between functions without copying it.
	// 2. Modify the original value in a function.

	some_int := 42          // Declare an int variable
	var p *int              // Declare a pointer to an int without assigning it
	var p2 *int = &some_int // Declare a pointer to an int and assign it the address of anInt

	// fmt.Println("Value of p:", *p) // this will error because p is nil

	if !is_valid_pointer(p) {
		println("p is nil, it does not point to any int value")
	} else {
		println("Value of p:", *p) // This would print the value if p was initialized
	}

	if !is_valid_pointer(p2) {
		println("p2 is nil, it does not point to any int value")
	} else {
		println("Value of p2:", *p2)
	}

	*p2 = *p2 + 1
	println("Value of p2 after incrementing:", *p2)
	println("Value of some_int after incrementing through p2:", some_int)
}
