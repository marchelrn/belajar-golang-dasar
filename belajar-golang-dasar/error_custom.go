package main

import "fmt"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func saveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Validation Error"}
	}
	if id != "Marchel" {
		return &notFoundError{"Data Not Found"}
	}
	return nil
}

func main() {
	err := saveData("Marche", nil)
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Validation Error : ", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("NotFound Error : ", notFoundErr.Error())
		} else {
			fmt.Println("Unknown Error : ", err.Error())
		}
	} else {
		fmt.Println("Saved Data")
	}
}
