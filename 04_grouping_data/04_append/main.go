package main

import "fmt"

func main() {

	// Composite literal syntax to create a new Slice []type{values}.
	// Slices are backed by an array implementation
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// the append built in function takes a slice and appends to it a variadic amount of value(s)
	// appends a single value to the end of x
	x = append(x, 77)
	fmt.Println(x)

	// appends multiple values to the end of x
	x = append(x, 100, 202, 1001)
	fmt.Println(x)

	y := []int{0, 2000, 2222, 1234, 666, 42}

	// Appends all the values from y to the end of x. Since append second argument is of type ...T
	// The values of slice y need to be unfurled by using the ... operation
	// x = append(x, y) results in: "cannot use y (type []int) as type int in append" error
	x = append(x, y...)

	fmt.Println(x)
}
