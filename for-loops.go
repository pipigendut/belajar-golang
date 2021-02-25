package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke ", counter2)
	}

	slice := []string{"Akbar", "Irma", "Menikah"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// with index
	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	// without index
	for _, value := range slice {
		fmt.Println(value)
	}

	person := map[string]string{
		"name":    "Akbar",
		"address": "Bekasi",
	}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
