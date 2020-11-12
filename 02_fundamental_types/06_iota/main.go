package main

import (
	"fmt"
)

// An iota increments by 1 every iota
const (
	a = 2000
	b = 2000 + iota
	c = 2000 + iota
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}