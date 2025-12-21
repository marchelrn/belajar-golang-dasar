package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"false" max:"10"`
}

func readField(value any) string {
	valueType := reflect.TypeOf(value)
	fmt.Println("Value Type : ", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))

	}
	return reflect.TypeOf(value).String()
}

func isValid(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			if data == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	//readField(Sample{"Marchello"})
	//readField(Person{"Marchel", "", ""})
	person := Person{
		Name:    "ada",
		Email:   "",
		Address: "dad",
	}
	fmt.Println(isValid(person))
}
