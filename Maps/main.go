package main

import (
	"fmt"
)

func main() {
	// // In "Maps" all the keys must be of same type and all the values must be of same type

	// using "strings" as key types in the map
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through maps
	for key, value := range menu {
		fmt.Printf("%v : %v\n", key, value)
	}

	// using "ints" as key types in the map
	phoneBook := map[int]string{
		97325: "Maa",
		97326: "Baba",
		97327: "Dada",
		97328: "Ami",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[97325])

	// looping through maps
	for key, value := range phoneBook {
		fmt.Printf("%v : %v\n", key, value)
	}

	// updating a value in a map
	phoneBook[97328] = "Shouvik"
	for key, value := range phoneBook {
		fmt.Printf("%v : %v\n", key, value)
	}
}
