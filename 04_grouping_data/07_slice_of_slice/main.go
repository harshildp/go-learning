package main

import "fmt"

func main() {

	a := []string{"Texas", "Austin"}
	b := []string{"New York", "Albany"}

	// A slice of a slice is denoted by [][]type
	x := [][]string{a, b}

	fmt.Println(x)

	// Looping over the values of a slice of a slice requires a nested for
	for _, v := range x {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
