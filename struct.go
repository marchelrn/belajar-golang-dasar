package main

import "fmt"

func main() {
	customer := customerInfo(0)

	customer1 := Customer{
		Name:    "Joko",
		Address: "Solo",
		Age:     30,
	}

	customer.sayHello("Eko")
	customer1.sayHello("doni")

}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name, "My Address is", customer.Address, "My Age is", customer.Age)
}

type Customer struct {
	Name    string
	Address string
	Age     int
}

func customerInfo(n int) Customer {
	customers := []Customer{
		{Name: "Marchel", Address: "Manado", Age: 19},
		{Name: "Budi", Address: "Jakarta", Age: 20},
		{Name: "Siti", Address: "Bandung", Age: 21},
	}
	if n < 0 || n > len(customers) {
		return Customer{Name: "NA", Address: "NA", Age: 0}
	}
	return customers[n]
}
