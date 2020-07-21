package main

import "fmt"

func main() {
	// For loops
	i := 1

	for i <= 10 {
		fmt.Println(i)

		// Shorthand for i = i + 1
		i++
	}

	// Relational Operators include ==, !=, <, >, <=, and >=

	// A for loop can be defines also like this, but semicolons are needed
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
