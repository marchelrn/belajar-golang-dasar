package main

import "fmt"

func main() {

	name := "Marchel"

	if name == "Marchel" {

		fmt.Println("Hallo, Marchel")
	} else if name == "Doni" {
		fmt.Println("Hallo, Doni")
	} else {
		fmt.Println("Hai, Boleh Kenalan?")
	}

	if length := len(name); length >= 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Mantap")
	}

}
