package main

import "fmt"

func main() {

	// Composite literal syntax to create a new Slice []type{values}.
	// Slices are backed by an array implementation
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// : operator can be used to slice a slice. x[startIndex : endIndex]
	// The startIndex is inclusive in the new slice whereas the endIndex is exclusive

	// Indexes 5 to the end
	y := x[5:]

	// Indexes from the start up till before index 3
	z := x[:3]

	// Indexes 2 up till before index 7
	a := x[2:7]

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
