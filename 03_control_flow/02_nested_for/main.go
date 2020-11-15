package main

import (
	"fmt"
)

func main() {

	// for every letter in the alphabet {}
	for i := 1; i <= 26; i++ {
		fmt.Println(i)

		// print three times
		for j := 0; j < 3; j++ {

			// %#U is the unicode output for a value
			fmt.Printf("\t%#U\n", i + 64)
		}
	}
}