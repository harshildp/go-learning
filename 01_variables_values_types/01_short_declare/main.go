package main

import "fmt"

// func main inside package main is the entry point to an application in Golang
func main() {

	// := declares and assigns a value to an identifier
	// all declared variables must be used or else Go compiler
	// will throw an unused error
	x := 42
	y := "Konnichiwa"
	z := true

	// Print outputs to the console and uses a variadic parameter ...interface{} 
	// i.e. it can take N number of arguments of varying types as all types are
	// of type interface{} at there base
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
}