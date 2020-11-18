package main

import "fmt"

// A struct is a composite data type which can contain fields of many other types
// It can be declare just like a new type is declare
type person struct {
	first string
	last  string
}

func main() {

	// A struct can be assigned to a variable as follows
	p1 := person{
		first: "Mace",
		last:  "Windu",
	}

	fmt.Println(p1)

	// The individual fields in a struct can be accessed like so
	fmt.Println(p1.first, p1.last)
}
