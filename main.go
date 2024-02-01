package main

import "fmt"

func updateName(n string) {
	n = "benjamin nkem"
}

func main() {

	name := "nkem"

	updateName(name)

	fmt.Println(name)

}
