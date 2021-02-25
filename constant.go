package main

import "fmt"

func main() {
	// variable const tidak bisa diubah lagi nilai-nya
	// tidak error jika tidak dipake
	const firstName string = "Akbar"
	const lastName = "Maulana"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// deklarasi multiple const
	const (
		wifeFirstName string = "Irma"
		wifeLastName         = "Yusnita"
	)
}
