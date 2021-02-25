package main

import "fmt"

func main() {
	var (
		names = [...]string{
			"Akbar", "Irma", "Menikah", "Tahun", "Depan",
		}
		sliceLowHigh = names[0:2]
		sliceLow     = names[2:]
		sliceHigh    = names[:4]
		sliceAll     = names[:]
	)

	fmt.Printf("%s\n", sliceLowHigh)
	fmt.Printf("%s\n", sliceLow)
	fmt.Printf("%s\n", sliceHigh)
	fmt.Printf("%s\n", sliceAll)
	fmt.Printf("Cap %d\n", cap(sliceHigh))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[4:6]
	daysAppend1 := append(daysSlice1, "Libur")
	fmt.Println(daysAppend1)
	fmt.Println(days)

	makeSlice := make([]string, 4, 5)
	makeSlice[0] = "Akbar"
	makeSlice[1] = "Maulana"
	makeSlice[3] = "Maulana"
	fmt.Println("makeSlice", makeSlice)

	copySlice := make([]string, len(makeSlice), cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	fmt.Printf("iniArray %d\n", iniArray)
	fmt.Printf("iniSlice %d\n", iniSlice)
}
