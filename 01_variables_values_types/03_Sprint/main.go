package main

import (
	"fmt"
)

// Declares a variable with an identifer ex. "x" with a type 
// Assigns a value
var x int = 42
var y string = "Konnichiwa"
var z bool = false

func main() {

	// Sprint is used to print to a String
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Print(s)
}