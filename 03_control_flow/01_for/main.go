package main

import (
	"fmt"
)

// Control flow is the order in which an application is executed.
// Generally applications execute sequentially. But there are also
// iterative and conditional blocks for executing code in a
// non-sequential manner. The for, if and switch blocks are used to do this

func main() {

	// for var, condition, post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
