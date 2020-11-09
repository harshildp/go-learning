package main

import (
	"fmt"
)

// Can create a new type by providing an Identifier and an underlying Type
type newtype int 
var x newtype

func main() {
	// The type of variable x is newtype not the underlying type int
	fmt.Printf("%v\t%T\n", x, x)
	x = 77
	fmt.Printf("%v\t%T", x, x)
}