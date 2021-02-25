package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := "asem", 2
	fmt.Println(c)
	fmt.Println(d)
}
