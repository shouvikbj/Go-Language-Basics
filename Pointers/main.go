package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func updateNameUisngPointer(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"
	updateName(name)
	fmt.Println(name)

	fmt.Println("Memory address of name is:", &name)
	m := &name
	fmt.Println("Memory address of name is:", m)
	fmt.Println("Value at memory address:", *m)

	updateNameUisngPointer(m)
	fmt.Println(name)
}
