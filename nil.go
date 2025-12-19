package main

import "fmt"

func main() {
	data := newMap("Marchel")
	if data == nil {
		fmt.Println("data kosong")
	}
	fmt.Println(data["Name"])
}

func newMap(name string) map[string]string {
	if name == " " {
		return nil
	}
	return map[string]string{
		"Name": name,
	}
}
