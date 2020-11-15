package main

import (
	"fmt"
)

func main() {
	i := 2000

	// for {} while keep looping endlessly. A break condition is required to exit the loop
	for {
		if i > 2020 {
			break
		}
		fmt.Println(i)
		i++
	}
}