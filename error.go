package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := Divide(100, 0)
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}
	fmt.Println("Hasil :", hasil)
}
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	return a / b, nil
}
