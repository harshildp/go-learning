package main

import (
	"fmt"
)

func main() {
	i := 77
	fmt.Printf("%d\t%b\t\t%#x\n", i, i, i)

	// bitwise operator to shift all the bits to the left 1 digit
	j := i << 1
	fmt.Printf("%d\t%b\t%#x", j, j, j)
}