package main

import "fmt"

type Addresses struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Addresses) {
	address.Country = "Indonesia"
}

func main() {
	address := Addresses{"Bandung", "Jawa Barat", "Malaysia"}
	changeCountryToIndonesia(&address)
	fmt.Println(address)
}
