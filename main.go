package main

import (
	"fmt"
)

func main() {

	age := 18
	name := "Benjamin"
	points := 4.5

	//  Print
	fmt.Print("Hello, ")
	fmt.Print("World \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("Some random string")
	fmt.Println("anther random string")

	fmt.Println("My age is", age, "and my name is", name)

	// Printf (Formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v\n", age, name)
	fmt.Printf("My age is %q and my name is %q\n", age, name) // does not work for int
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("type: score is %T\n", points)
	fmt.Printf("you scored %0.1f points!\n", points)
	fmt.Printf("you scored %T points!\n", points)
	fmt.Printf("binary of %d is %b\n", age, age) // convert to binary

	// Sprintf (save formatted string)
	formStr := fmt.Sprintf("My age is %v and my name is %v\n", age, name)

	fmt.Println(formStr)
}
