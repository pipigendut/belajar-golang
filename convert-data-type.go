package main

import "fmt"

func main() {
	// number
	var (
		nilai32 int32 = 12345
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai64)
	)

	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// string
	var (
		firstName      = "Akbar"
		e         byte = firstName[0]
		eString        = string(e)
	)

	fmt.Println(eString)
}
