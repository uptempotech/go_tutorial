package main

import "fmt"

func main() {

	// An Array holds a fixed number of values of the same type
	var favNums2 [5]float64

	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.61

	// Access to the value is done by supplying the index number
	fmt.Println(favNums2[3])

	// Another way of initializing an array
	favNums3 := [5]float64{
		1,
		2,
		3,
		4,
		5,
	}

	// How to iterate through an array (Use _ if a value isn't used)
	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

	// Slices are like arrays but you leave out the size
	numSlice := []int{
		5,
		4,
		3,
		2,
		1,
	}

	// A slice can be created by defining the first index value to
	// take through the last

	// numSlice2 == [2,1]
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice[0] = ", numSlice2[0])

	// If the first index is not specified, it defaults to 0
	// if the last index is not specified, it defaults to max
	fmt.Println("numSlice[:2]", numSlice[:2])
	fmt.Println("numSlice[2:]", numSlice[2:])

	// To create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10)

	// A slice can be copied from one to another
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0])

	// Append values to the end of a slice
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6])
}
