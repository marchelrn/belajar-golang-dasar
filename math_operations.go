package main

import "fmt"

func main() {
	var a = 10
	var b = 2
	c := a + b

	fmt.Println(c)

	//augmented assign

	c += 10
	fmt.Println(c)

	c %= 12
	fmt.Println(c)

	value := 10

	value %= 10
	fmt.Println(value)

	//unary operator
	value++
	fmt.Println(value)

	value--
	fmt.Println(value)
}
