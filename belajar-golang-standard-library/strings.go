package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Marchel Manullang", "chel"))
	fmt.Println(strings.Split("Marchel Manullang", "  "))
	fmt.Println(strings.ToLower("Marchel Manullang"))
	fmt.Println(strings.ToUpper("marchel manullang"))
	fmt.Println(strings.Trim("  Marchel Manullang  ", " "))
	fmt.Println(strings.ReplaceAll("Marchel Manullang", "Marchel", "Achel"))
}
