package main

import "fmt"

func main() {
	counter := 0
	counter2 := 0
	increment := func() {
		counter++ // ini menyembabkan variable counter berubah(closure)
	}

	increment2 := func() {
		counter2 := 0 // dengan mendlekarasikan ulang variable counter2 didalam scope anonymouse func, counter2 diluar func ini tidak akan berubah
		counter2++
	}

	increment()
	increment2()

	fmt.Println("counter = ", counter)
	fmt.Println("counter2 = ", counter2)
}
