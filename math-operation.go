package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 20
		c = a * b
	)
	fmt.Println(c)

	// augmented assignments
	var i = 10
	i += 10 // i = i + 10
	i -= 10 // i = i - 10
	i *= 10 // i = i * 10
	i %= 10 // i = i % 10
	fmt.Println(i)

	// unary operation
	var j = 10
	j++ // j = j + 1
	j-- // j = j - 1
	fmt.Println(j)
}
