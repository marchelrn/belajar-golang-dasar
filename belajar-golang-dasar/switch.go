package main

import (
	"fmt"
)

func main() {

	name := ""

	switch name {

	case "Marchel":
		fmt.Println("Halo, Marchel")
	case "Doni":
		fmt.Println("Halo, Doni")
	case "Holi":
		fmt.Println("Halo, Holi")
	default:
		fmt.Println("Hai, Boleh kenalan?")

	}

	// Switch dengan short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Halo, Selamat datang")
	}

	// Switch tanda kondisi

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length <= 5:
		fmt.Println("Halo, Selamat datang")
	default:
		fmt.Println("Nama tidak valid")
	}
}
