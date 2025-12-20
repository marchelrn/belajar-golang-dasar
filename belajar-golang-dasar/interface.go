package main

import "fmt"

func main() {
	person := Person{Name: "Marchel"}
	getHello(person)

	animal := Animal{Age: 5}
	getHello(animal)
}

type hasName interface {
	getName() string
}

func getHello(value hasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Age int
}

func (animal Animal) getName() string {
	return fmt.Sprintf("Meow ur age %d", animal.Age)
}
