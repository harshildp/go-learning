package main

import "fmt"

func main() {

	// Structs can be created anonymously like below. Here there is no identifier associated with the struct.
	// It can be utilized when the scope required for the struct is limited and it may be beneficial to
	// not create a new type to hold the struct
	p1 := struct {
		first    string
		last     string
		episodes []int
		design   map[string]string
	}{
		first:    "Mace",
		last:     "Windu",
		episodes: []int{1, 2, 3},
		design: map[string]string{
			"race":            "human",
			"lightsaberColor": "green",
		},
	}

	fmt.Println(p1)
}
