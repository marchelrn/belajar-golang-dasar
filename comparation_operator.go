package main

import "fmt"

func main() {

	var name1 = "Marchel"
	var name2 = "Marchel"

	result := name1 == name2
	fmt.Println(result)

	value1 := 100
	value2 := 200

	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)

}
