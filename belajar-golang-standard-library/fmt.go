package main

import "fmt"

func main() {
	firstName := "Marchel"
	lastName := "Manullang"
	var net float32 = 1000000

	discount := 0.1
	total := net - (net * float32(discount))
	fmt.Printf("Helllo , %s %s  %s %.2f%% \n", firstName, lastName, "You got a discount of ", discount*100)
	fmt.Printf("Total : %v \n", total)
}
