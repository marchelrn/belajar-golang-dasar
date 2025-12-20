package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("true")
	if err == nil {
		println(parseBool)
	} else {
		println("Error:", err.Error())
	}

	parseInt, err := strconv.Atoi("1000")
	if err == nil {
		println(parseInt)
	} else {
		println("Error:", err.Error())
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	intToString := strconv.Itoa(1000)
	fmt.Println(intToString)
}
