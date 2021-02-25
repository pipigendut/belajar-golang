package main

import "fmt"

func main() {
	// membuat type data baru atau alias dari tipe data yang sudah ada
	type typeNoKTP string // membuat tipe data alias string menjadi noKTP
	type typeTelp int32

	var (
		noKtpAkbar  typeNoKTP = "3216060"
		noKtpIrma   typeNoKTP
		noTelpAKbar typeTelp = 21222
	)
	noKtpIrma = "3216060"

	fmt.Println(noKtpAkbar)
	fmt.Println(noKtpIrma)
	fmt.Println(noTelpAKbar)
}
