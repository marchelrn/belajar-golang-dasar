package main

import "fmt"

func random() any {
	return "test"
}

func main() {

	result := random()
	//resultString := result.(string)
	//fmt.Printf("Type %T\n", resultString)
	//fmt.Printf("Type %T\n", result)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	case func():
		fmt.Println("Function", value)
	default:
		fmt.Println("Unknown Type")
	}

}
