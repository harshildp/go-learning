package main

import "fmt"

func main() {

	// map can be created using map[keyType]valueType syntax
	// a map stores keys that have values associated to them
	x := map[string][]string{
		"Mace Windu": []string{"Purple", "Jedi"},
		"Yoda":       []string{"Green", "Jedi"},
		"Darth Maul": []string{"Red", "Sith"},
	}

	fmt.Println(x)

	// The type of x is map[string][]string
	fmt.Printf("%T\n", x)

	// Look up from the map based on a given key
	fmt.Println(x["Yoda"])

	// Loop through all the keys and their values
	for k, val := range x {
		fmt.Println("Name is", k)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

	// Adding a new record to the map. Use a key and assign its value
	x["Obi Wan Kenobi"] = []string{"Blue", "Jedi"}

	for k, val := range x {
		fmt.Println("Name is", k)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

	// A key which doesn't exist in the map returns the default value for the value type field
	// [] in this case for value type []string
	fmt.Println(x["Ahsoka Tano"])

	// The following syntax can be used to check for existence of a key in a map
	// v will be the value assigned to that key (default value if key doesn't exist)
	// ok will be a bool value denoting whether the key was found in the map or not
	v, ok := x["Ahsoka Tano"]

	// the ok variable can then be utilzied to perform logic if the key existed. Like print its value
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key wasn't found in the map")
	}

	// The delete function can be used to delete a specific key and its value from a map
	// delete(mapName, keyToBeDeleted)
	delete(x, "Darth Maul")

	for k, val := range x {
		fmt.Println("Name is", k)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

	// deleting a key which already doesn't exist, does not throw an error
	delete(x, "Darth Maul")
}
