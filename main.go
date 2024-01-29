package main

import (
	"fmt"
	"math"
)

func somethingElse() {
	fmt.Println("this is ran from the 'something else' function.")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func calculateCircumference(radius float64) float64 {
	return 2 * math.Pi * radius
}

func main() {
	// var firstName string = "Benjamin"
	// lastName := "Nkem"

	// fmt.Println(firstName, lastName)

	// var age int8 = 100
	// fmt.Println(age)

	// gpa := 0.5
	// fmt.Println(gpa)

	// somethingElse()
	// greet("Benjamin Nkem")

	// result := calculateCircumference(30)
	// fmt.Println(result)

	var ageOne int8 = 20

	// uint - you can't have a negative number
	var num uint8 = 30

	fmt.Println(ageOne, num)

}
