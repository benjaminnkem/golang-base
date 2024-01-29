package main

import (
	"fmt"
)

func main() {

	// var ages [3]int = [3]int{1, 2, 3} // has a fixed length
	var ages = [3]int{1, 2, 3}

	names := [4]string{"benjamin", "tochi", "junior", "nkem"}

	names[1] = "francis"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 30

	scores = append(scores, 34)

	fmt.Println(scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "cooper")

	fmt.Println(rangeOne)

}
