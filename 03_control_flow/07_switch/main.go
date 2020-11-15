package main

import (
	"fmt"
)

func main() {

	// Conditionless switch statement. Executes the first true statement
	switch {
	case false:
		fmt.Println("Doesn't print")

	// case can be any expression which can be evaluated to a bool
	case 2 == 2:
		fmt.Println("Prints")
	case true:
		fmt.Println("Prints2, doesn't actually print")
	}
}
