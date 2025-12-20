package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{}

	// person["name"] = "Marchel"
	// person["address"] = "Manado"

	person := map[string]string{
		"name":    "Marchel",
		"address": "Manado",
	}

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Marchel"
	book["ups"] = "Salah"

	delete(book, "ups")
	fmt.Println(book)

}
