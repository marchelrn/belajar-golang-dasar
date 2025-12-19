package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// copy by value
	//address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//address2 := address1
	//
	//address2.City = "Bandung"
	//address2.Country = "Indonesia"
	//address2.Province = "Jawa Barat"
	//
	//fmt.Println(address1)
	//fmt.Println(address2)

	// pointer
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"
	address2.Country = "Indonesia"
	address2.Province = "Jawa Barat"

	fmt.Println(address1)
	fmt.Println(address2)
}
