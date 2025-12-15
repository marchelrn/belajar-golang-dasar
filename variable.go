package main

import "fmt"

func main() {

	name := "Marchel"
	fmt.Println(name)

	name = "Manullang"
	fmt.Println(name)

	fmt.Println(name)
}

func callName() {
	var namaDepan string
	var namaBelakang string
	var umur int
	var fullName string

	fmt.Println("Enter First Name, Last Name, Age:")

	fmt.Scan(&namaDepan)
	fmt.Scan(&namaBelakang)
	fmt.Scan(&umur)

	fullName = namaDepan + " " + namaBelakang

	fmt.Println(fullName)
	fmt.Println(umur)
}
