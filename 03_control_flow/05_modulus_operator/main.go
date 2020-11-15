package main

import (
	"fmt"
)

func main() {

	// for var, condition, post {}
	for i := 0; i <= 100; i++ {

		// The % operator outputs the remainder of a division.
		// For example 3 % 2 = 1 because 3 / 2 has a remainder of 1
		fmt.Printf("Remainder of %v / 4 is %v\n", i, i % 4)
	}
}