package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke -", counter)
	}
	fmt.Println("Done")

	names := []string{"Marchel", "Matthew", "Manullang"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for idx, name := range names {

		fmt.Println("index", idx, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
