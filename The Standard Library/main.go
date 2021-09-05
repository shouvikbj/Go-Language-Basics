package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// // "strings" package

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "ninjas"))

	fmt.Println(strings.ReplaceAll(greeting, "friends", "ninjas"))

	fmt.Println(strings.ToUpper(greeting))

	fmt.Println(strings.Index(greeting, "th"))

	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("Original string value =>", greeting)

	// // "sort" package

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // alters the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // finds the index of 30 in the sorted slice
	fmt.Println(index)

	index = sort.SearchInts(ages, 90) // return (lastIndex+1) for elements if not found in the slice
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "luigi"))

}
