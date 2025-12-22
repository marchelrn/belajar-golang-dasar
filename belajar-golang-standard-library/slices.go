package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"Marchel", "Manullang", "Joko", "Widodo"}
	value := []int{10, 20, 30, 40}
	fmt.Println(slices.Min(name))
	fmt.Println(slices.Max(value))
	fmt.Println(slices.Contains(name, "Marchel"))
	fmt.Println(slices.Contains(value, 20))
}
