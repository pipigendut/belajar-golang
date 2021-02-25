package main

import "fmt"

func main() {
	// array of string
	var names [3]string
	names[0] = "Akbar"
	names[1] = "Irma"
	names[2] = "Menikah"

	fmt.Println(names)

	// array of int
	var scores [3]int
	scores[0] = 1
	fmt.Println(scores)

	var values = [3]int{1, 3, 4}
	fmt.Println(values)
}
