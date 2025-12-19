package main

import "fmt"

func Ups() any {
	return "ups"
}

func main() {
	ups := Ups()
	fmt.Println(ups)
}
