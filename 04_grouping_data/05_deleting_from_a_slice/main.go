package main

import "fmt"

func main() {

	// Composite literal syntax to create a new Slice []type{values}.
	// Slices are backed by an array implementation
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// the append function along with slicing can used to delete values from a slice
	// Deletes the value 3 (index position 2) from the slice x
	x = append(x[:2], x[3:]...)

	fmt.Println(x)
}
