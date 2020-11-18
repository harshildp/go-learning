package main

import "fmt"

type person struct {
	first string
	last  string
}

// A struct can be embedded with another struct
// Here the first field of the jedi struct is a person struct
// Essentially the jedi class is building on the person struct while reusing
// the person structs fields
type jedi struct {
	person
	color string
}

func main() {
	p1 := person{
		first: "Mace",
		last:  "Windu",
	}

	// The jedi struct can be assigned the following ways
	// Referencing an already assigned person p1
	j1 := jedi{
		person: p1,
		color:  "purple",
	}

	// Creating a new person as the time of assigning j2
	j2 := jedi{
		person: person{
			first: "Yoda",
			last:  "",
		},
		color: "green",
	}

	fmt.Println(j1)
	fmt.Println(j2)

	// The fields of the person struct are promoted directly to the jedi struct.
	// Therefore they can be accessed directly by using jedi.fieldName
	fmt.Println(j1.first, j1.last, j1.color)

	// They can still be accessed using the full path ex. jedi.person.fieldName
	fmt.Println(j2.person.first, j2.person.last, j1.color)

	// This enables you to reference individual fields in case some of the names overlap
}
