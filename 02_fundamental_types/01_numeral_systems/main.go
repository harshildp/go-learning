package main

import (
	"fmt"
)

func main() {
	i := 77

	// %d prints a digit in decimal format
	// %b prints in binary format
	// %#x prints in hexadecimal format
	fmt.Printf("%d\t%b\t%#x", i, i, i)
}