package main

import "fmt"

func main() {

	//Break

	for i := 0; i <= 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke - ", i)
	}

	//Continue

	for i := 0; i <= 100; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println("ini adalah bilangan ganjil", i)
	}
}
