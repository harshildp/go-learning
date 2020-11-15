package main

import (
	"fmt"
)

func main() {

	x := 42

	// Switch statement with a variable to use for comparisons.
	// Executes the first true statement
	switch x {
	case 2:
		fmt.Println("Just the number 2")
	case 41:
		fmt.Println("Number 41")
	case 42:
		fmt.Println("You have found the meaning of life")
	case 66:
		fmt.Println("Execute Order 66")
	case 77:
		fmt.Println("Just another nice number")
	}
}
