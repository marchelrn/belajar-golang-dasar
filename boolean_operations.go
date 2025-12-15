package main

import "fmt"

func main() {
	umur := 18
	gender := "Laki-laki"

	if umur >= 19 && gender == "Perempuan" {
		fmt.Println("Kamu bisa masuk")
	} else if umur >= 19 || gender == "Laki-laki" {
		fmt.Println("Kamu bisa masuk")
	} else {
		fmt.Println("Kamu tidak bisa masuk")
	}

	result := 80
	absent := 3

	var succeedTest bool = result >= 80
	var goalAbsent bool = absent < 2

	var succeed bool = succeedTest == false && goalAbsent == false

	fmt.Println(succeed)
}
