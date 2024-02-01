package main

import "fmt"

func main() {

	menu := map[string]float64{
		"pizza":  1.11,
		"garri":  20.99,
		"rice":   4.99,
		"tomato": 2.99,
	}

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

}
