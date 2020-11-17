package main

import "fmt"

func main() {

	// Composite literal syntax to create a new Slice []type{values}.
	// Slices are backed by an array implementation
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Slices are 0 index based and can be reassigned like so
	x[3] = 77

	// A for-each loop to loop through each index and value in the range of x
	for i, v := range x {
		fmt.Printf("Value at index %v is %v\n", i, v)
	}

	// x is of type []int.s
	fmt.Printf("%T", x)
}
