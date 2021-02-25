package main

import "fmt"

func multiple(x string, y string, z int) (string, string, int) {
	return x, y, z
}

func main() {
	a, b, c := multiple("Akbar", "Maulana", 100)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, e, _ := multiple("Akbar", "Maulana", 100)
	fmt.Println(e)
}
