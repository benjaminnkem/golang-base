package main

import "fmt"

func main() {

	myBill := newBill("benjamin")

	myBill.addItem("spaghetti", 4.95)
	myBill.tip = 3

	fmt.Println(myBill.format())
	// fmt.Println(myBill)

}
