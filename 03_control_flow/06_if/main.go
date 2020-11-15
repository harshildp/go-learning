package main

import (
	"fmt"
)

func main() {
	x := 42

	// if condition {} is true it runs the code block
	if x == 42 {
		fmt.Println("This is the meaning of life")

		// else catches all the cases which haven't entered an if block
	} else {
		fmt.Println("This is just a number")
	}
}
