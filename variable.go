package main

import "fmt"

func main() {
	var name string

	name = "Akbar Maulana"
	fmt.Println(name)

	name = "Irma Yusnita"
	fmt.Println(name)

	var wifeName = "Irma Yusnita"
	fmt.Println(wifeName)

	var loveName string = "Irma Yusnita"
	fmt.Println(loveName)

	var int32 int
	int32 = 30
	fmt.Println(int32)

	var int64 int64 = 64
	fmt.Println(int64)

	// ':=' untuk menggantikan 'var'
	country := "Indonesia"
	fmt.Println(country)
	country = "Amerika"
	fmt.Println(country)

	score := 64
	fmt.Println(score - int32)

	// multiple deklarasi variable
	var (
		firstName = "Akbar"
		lastName  = "Maulana"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
