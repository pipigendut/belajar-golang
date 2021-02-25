package main

import "fmt"

func main() {
	name := "Akbar"

	switch name {
	case "Akbar":
		fmt.Println("Hallo Akbar")
	case "Irma":
		fmt.Println("Hallo Irma")
	default:
		fmt.Println("Hallo Guest")
	}

	// short statement
	switch lengthName := len(name); lengthName > 4 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Terlalu pendek")
	}

	// switch without expresion
	var lengthName int
	lengthName = len(name)
	switch {
	case lengthName > 5:
		fmt.Println("Terlalu panjang", lengthName)
	case lengthName < 5:
		fmt.Println("Terlalu pendek", lengthName)
	default:
		fmt.Println("Sempurna", lengthName)
	}
}
