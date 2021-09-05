package main

import "fmt"

func main() {
	// Arrays

	var ages1 [3]int = [3]int{20, 25, 30}
	var ages2 = [3]int{10, 15, 20}

	names := [4]string{"shouvik", "soumen", "piku", "kittu"}

	fmt.Println(ages1, len(ages1))
	fmt.Println(ages2, len(ages2))
	fmt.Println(names, len(names))

	names[3] = "dada"
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	rangeFour := names[:]

	fmt.Println(rangeOne, rangeTwo, rangeThree, rangeFour)

	rangeOne = append(rangeOne, "name 4")
	fmt.Println(rangeOne)

}
