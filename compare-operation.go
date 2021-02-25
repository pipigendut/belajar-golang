package main

import "fmt"

func main() {
	var (
		name1        = "Akbar"
		name2        = "Akbar"
		value1       = 100
		value2       = 200
		result1 bool = name1 == name2
		result2 bool = value1 < value2
	)

	fmt.Println(result1)
	fmt.Println(result2)
}
