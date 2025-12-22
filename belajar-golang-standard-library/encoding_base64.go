package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

func main() {
	value := []byte("101010011")

	encoded := base64.StdEncoding.EncodeToString(value)
	fmt.Println(encoded)

	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println(string(decode))
	tipe := reflect.TypeOf(encoded)
	fmt.Println(tipe)

}
