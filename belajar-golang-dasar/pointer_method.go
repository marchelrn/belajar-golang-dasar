package main

import "fmt"

type Cust struct {
	Name   string
	Age    int
	Gender string
	Status bool
}

func (cust *Cust) married() bool {
	if cust.Gender == "Men" && cust.Status == true {
		return true
	} else if cust.Gender == "Women" && cust.Status == true {
		return true
	}
	return false
}

func (cust *Cust) greet() string {
	married := cust.married()
	if married == true {
		return "Hello Mr/Mrs " + cust.Name
	}
	return "Hello Ms " + cust.Name
}

func main() {
	customer := []Cust{
		{Name: "John",
			Age:    30,
			Gender: "Men",
			Status: true,
		},
		{Name: "Jane",
			Age:    25,
			Gender: "Women",
			Status: false,
		},
	}
	cust_data := customer[1].greet()
	fmt.Println(cust_data)
}
