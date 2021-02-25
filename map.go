package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Akbar",
		"address": "Bekasi",
	}
	person["title"] = "Programmer"

	fmt.Println("person", person)
	fmt.Println("personName", person["name"])
	fmt.Println("personTitle", person["title"])

	book := make(map[string]string)
	book["author"] = "Akbar"
	book["title"] = "Cara menjadi Golang Developer"
	fmt.Println("book", book)
	fmt.Println("bookAuthor", book["author"])

	delete(book, "title") // delete map depends key
	fmt.Println("newBook", book)
}
