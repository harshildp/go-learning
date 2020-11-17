package main

import "fmt"

func main() {

	// Composite literal syntax to create a new array [capacity]type{values}
	x := [5]int{1, 2, 3, 4, 5}

	// Arrays are 0 index based and can be reassigned like so
	x[3] = 77

	// A for-each loop to loop through each index and value in the range of x
	for i, v := range x {
		fmt.Printf("Value at index %v is %v\n", i, v)
	}

	// x is of type [5]int. The length is part of the type in Go
	fmt.Printf("%T", x)
}
