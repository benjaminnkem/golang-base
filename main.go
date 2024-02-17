package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Created bill -", b.name)

	return b
}

func main() {

	// myBill := createBill()
	// fmt.Println(myBill)

	ben := createUser("benjamin", 40)

	fmt.Print(ben)

}
