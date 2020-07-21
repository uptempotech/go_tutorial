package main

import "fmt"

func main() {
	// declaration of variable x
	var x int = 5
	fmt.Println("x = ", x)

	// declaration of multiple variables
	var (
		a = 10
		b = 15
	)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("a = ", a, "b = ", b)

	// shorthand declaration of variables
	y := 7
	fmt.Println("y = ", y)

	// declaration of constants
	const pi float64 = 3.14272345

	var name string = "Joe Tester"
	fmt.Println("name = ", name)
}
