package main

import "fmt"

func sayHelloToPerson(key, value string) (iniKey, iniValue string) {
	iniKey = fmt.Sprintln("Key ini adalah", key)
	iniValue = fmt.Sprintln("Value ini adalah", value)

	return
}

func main() {
	var (
		person = map[string]string{
			"author": "Akbar",
			"wife":   "Irma",
		}
	)

	for key, value := range person {
		iniKey, iniValue := sayHelloToPerson(key, value)
		fmt.Println(iniKey)
		fmt.Println(iniValue)
	}
}
