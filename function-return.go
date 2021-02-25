package main

import "fmt"

func sayHelloTo(name string) string {
	fuu := fmt.Sprintf("Hello %v", name)
	return fuu
}

func main() {
	result := sayHelloTo("Akbar")
	fmt.Println(result)
}
