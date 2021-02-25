package main

import "fmt"

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRekursive(value int) int {
	// return Ternary(value == 1, 1, value*factorialRekursive(value-1)).(int)

	if value == 1 {
		return 1
	} else {
		return value * factorialRekursive(value-1)
	}
}

func main() {
	fmt.Println("Hasil = ", factorialLoop(5))
	fmt.Println("Hasil = ", factorialRekursive(5))
}
