package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func updateName2(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays,structs
	// // Non-Pointer Values
	name := "tifa"

	updateName(name)
	fmt.Println(name)

	name = updateName2(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	// // Pointer Wrapper Values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
