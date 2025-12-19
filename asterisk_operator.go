package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Manado", "Sulawesi Utara", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

	address1.City = "Bitung"
	fmt.Println(address1)
	fmt.Println(address2)

	address1 = Address{"Medan", "Sumatera Utara", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

	address2 = &address1
	address1.City = "Kotamobagu"
	fmt.Println(address1)
	fmt.Println(address2)

	address1 = Address{"Palembang", "Sumatera Selatan", "Indonesia"}
	address2 = &address1
	fmt.Println(address1)
	fmt.Println(address2)

	address2 = &Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	//address2.City = "Tomohon"
	//fmt.Println(address2)
}
