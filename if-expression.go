package main

import "fmt"

func main() {
	name := "irm"

	if name == "akbar" {
		fmt.Println("Hello", name)
	} else if name == "irma" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello Guest")
	}

	// short statement
	if lengthName := len(name); lengthName > 4 {
		fmt.Println("Terlalu panjang")
	} else if lengthName < 4 {
		fmt.Println("Terlalu pendek")
	} else {
		fmt.Println("Sempurna")
	}
}
