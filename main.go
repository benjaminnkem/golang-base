package main

import "fmt"

func updateName(n *string) {
	*n = "benjamin"
}

func updateNames(n []string) {
	n[0] = "maleek"
}

func main() {

	name := "nkem"

	m := &name

	updateName(m)

	// fmt.Println("variable 'name' memory location is:", m)
	// fmt.Printf("variable 'name' value at location %v is: %v\n", m, *m)

	names := []string{"benjamin", "collins", "david", "michael", "francisca"}

	updateNames(names)

	fmt.Println(name)
	fmt.Println(names)

}
