package main

import "strings"

func main() {
	hello := sayHelloWithFilter
	filter := filterName
	name := "anjIng"
	println(hello(name, filter))
}

func filterName(name string) string {
	if strings.EqualFold("anjing", "Anjing") {
		return "****"
	}
	return name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	return "Hello " + filter(name)
}
