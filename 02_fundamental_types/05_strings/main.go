package main

import (
	"fmt"
)

func main() {

	// " " double quotes can be used to create a string
	x := "Konnichiwa!"

	// ` ` backticks can be used to create a raw literal string
	y := `"Konnichiwa
	!"`

	fmt.Println(x)
	fmt.Println(y)
}