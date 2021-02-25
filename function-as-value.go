package main

import "fmt"

func getGoodBye(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	// func as value adalah assign function kesebuah variable
	sayGoodBye := getGoodBye
	result := sayGoodBye("AKbar")
	fmt.Println(result)
}
