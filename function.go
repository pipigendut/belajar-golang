package main

import "fmt"

func sayHello(i int) {
	fmt.Println("Hello World", i)
}

func sayBook(key string, value string) {
	fmt.Println(key, "=", value)
}

func main() {
	for i := 0; i <= 2; i++ {
		sayHello(i)
	}

	book := map[string]string{
		"Author": "Akbar",
		"Year":   "2020",
	}
	for key, value := range book {
		sayBook(key, value)
	}
}
