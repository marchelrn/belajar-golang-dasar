package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return validationError
	} else if id != "Marchel" {
		return notFoundError
	}
	return nil
}

func main() {
	err := getById("e")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error:", err)
		}
	}
}
