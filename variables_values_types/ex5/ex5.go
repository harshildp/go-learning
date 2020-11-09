package main

import (
	"fmt"
)

// Can create a new type by providing an Identifier and an underlying Type
type newtype int 
var x newtype
var y int

func main() {
	// The type of variable x is newtype not the underlying type int
	fmt.Printf("%v\t%T\n", x, x)
	x = 77
	fmt.Printf("%v\t%T\n", x, x)

	// Conversion, not casting, is used to convert from one type to another
	// Even tho underlying type of x is int
	// it cannot be directly assigned to y as x's actual type is newtype
	y = int(x)
	fmt.Printf("%v\t%T", y, y)

}