package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Marchel"
	names[1] = "Bintang"
	names[2] = "Rizky"
	names[3] = "Fadilla"

	fmt.Println(names[3])
	fmt.Println(len(names))

	var value = [4]int{
		1,
		3,
		4,
		5,
	}

	fmt.Println(value)

	value[1] = 30
	fmt.Println(value)
}
