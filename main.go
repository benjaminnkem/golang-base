package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(name string) {
	fmt.Println("Hi,", strings.ToUpper(name))
}

func sayBye(name string) {
	fmt.Println("Goodbye,", name)
}

func cycleNames(names []string, fn func(string)) {
	for _, name := range names {
		fn(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func calculateCircumference(radius float64) float64 {
	return 2 * math.Pi * float64(radius)
}

func findName(names []string, target string) string {
	for _, v := range names {
		if v == target {
			return "target found"
		}
	}

	return "not found"
}

func getInitials(name string) (string, string) {

	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""

}

func main() {

	names := []string{"benjamin", "michael", "collins", "king david", "francisca"}

	// cycleNames(names, sayGreeting)

	fmt.Printf("The circumference is: %0.3f to 3.d.p\n", calculateCircumference(26.5))

	// c1 := circleArea(26.5)
	// c2 := circleArea(40)

	// fmt.Printf("The area is: %0.3f to 3.d.p\n", c1)
	// fmt.Printf("The area is: %0.3f to 3.d.p\n", c2)

	result := findName(names, "collins")

	fmt.Println(result)

	name1, name2 := getInitials("benjamin nkem")

	fmt.Println(name1, name2)
}
