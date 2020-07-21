package main

import (
	"fmt"
)

// POINTERS

func main() {
	// Pass the value of a variale to the function
	x := 0
	changeVal(x)
	fmt.Println("x =", x)

	// If we pass a reference to the variale we can change the value
	// in a function
	changeXValNow(&x)
	fmt.Println("x =", x)

	// Get the address x point to with &
	fmt.Println("Memory Address for x =", &x)
}

func changeVal(x int) {
	// Has no effect on the value of x in main()
	x = 2
}

// * signals that we are being sent a reference to the value
func changeXValNow(x *int) {
	// Change the value at the memory address referenced by the pointer
	// * gives us access to the value the pointer points at

	// Store 2 in the memory address x refers to
	*x = 2
}
