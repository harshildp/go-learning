package main

import "fmt"

func main() {

	// Make function can be used to create a slice of a given length and capacity.
	// make(type, length, capacity)
	// If capacity is reached, Go will copy the values to a new underlying array with
	// double the capacity
	x := make([]string, 5, 5)
	fmt.Println(x)
	fmt.Println(len(x)) // prints length of x
	fmt.Println(cap(x)) // prints capacity of x

	x = []string{
		"Hi",
		"Hello",
		"Hey",
		"Sup",
		"Good Morning",
	}

	fmt.Println(x)

	// Adding another element to go above capacity 5 of x
	x = append(x, "Konnichiwa")
	fmt.Println(x)
	fmt.Println(len(x)) // length incremented to 6
	fmt.Println(cap(x)) // capacity doubled to 10

}
