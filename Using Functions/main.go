package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	sayGreeting("Mario")
	sayGreeting("Luigi")
	sayBye("Mario")
	sayBye("Luigi")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Printf("%0.2f %0.3f \n", a1, a2)

}
