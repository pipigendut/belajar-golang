package main

import "fmt"

type ParamFilter func(string) string

func sayHelloWithFilter(name string, filter ParamFilter) string {
	nameFiltered := filter(name)
	return fmt.Sprintf("Hello %s", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	result := sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println(result)

	otherResult := sayHelloWithFilter("Babi", spamFilter)
	fmt.Println(otherResult)
}
