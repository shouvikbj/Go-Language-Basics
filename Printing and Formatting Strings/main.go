package main

import "fmt"

func main() {

	age := 23
	name := "Shouvik Bajpayee"

	// Print()
	fmt.Print("Hello, ")
	fmt.Print("world!\n")
	fmt.Print("My name is ", name, " and my age is ", age, "\n")

	// Println()
	fmt.Println("Hello world!")
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("Name is of type %T \n", name)
	fmt.Printf("You scored %f points. \n", 225.55)
	fmt.Printf("You scored %0.3f points. \n", 225.55)

	// Sprintf (save formatted strings)
	var savedString = fmt.Sprintf("My name is %v and my age is %v \n", name, age)
	fmt.Println("The saved string is :", savedString)

}
