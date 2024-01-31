package main

import "fmt"

func main() {

	// i := 0

	// for i < 50 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 50; i++ {
	// 	fmt.Println(i)
	// }

	names := []string{"benjamin", "tochi", "junior", "nkem"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for index, value range names {
	// 	fmt.Println("the index is %v and value is %v", index, value)
	// }

}
