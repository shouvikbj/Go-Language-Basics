package main

import "fmt"

func main() {

	x := 0
	//  works like "while" loop
	for x < 5 {
		fmt.Println("Value of x is :", x)
		x++
	}

	// works like normal "for" loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is :", i)
	}

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// works like "forEach" loop
	// similar to how we iterate over "dictionaries" in "python"
	for index, value := range names {
		fmt.Printf("At index %v the name is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "new string" // doesn't update the actual slice
	}

	fmt.Println(names)

}
