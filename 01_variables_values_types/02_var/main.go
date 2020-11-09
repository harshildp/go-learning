package main

import (
	"fmt"
)

// Declares a variable with an identifier ex. "x" with a type 
// Assigns the zero_value associated with the type
// 0 for int, "" for String, false for bool, 0.0 for float
var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}