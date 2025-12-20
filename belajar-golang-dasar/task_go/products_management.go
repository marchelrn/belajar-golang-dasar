package main

import (
	"fmt"
)

func main() {
	result, found := getProductById(10)
	if !found {
		fmt.Println("Product not found")
		return
	}
	result.displayInfo()
	fmt.Printf("Status: %s\n\n", map[bool]string{true: "Available", false: "Out of Stock"}[result.isAvailable()])
}

type Product struct {
	Id       int
	Name     string
	Category string
	Price    float64
	Stock    int
}

func getProductById(id int) (Product, bool) {
	products := []Product{
		{Id: 1, Name: "Laptop", Category: "Electronics", Price: 1500.00, Stock: 10},
		{Id: 2, Name: "Smartphone", Category: "Electronics", Price: 800.00, Stock: 25},
		{Id: 3, Name: "Desk Chair", Category: "Furniture", Price: 120.00, Stock: 15},
		{Id: 4, Name: "Book", Category: "Stationery", Price: 20.00, Stock: 50},
		{Id: 5, Name: "Headphones", Category: "Electronics", Price: 150.00, Stock: 30},
	}
	for _, product := range products {
		if product.Id == id {
			return product, true
		}
	}
	return Product{Name: "Not Found", Stock: 0}, false
}

func (product Product) displayInfo() {
	fmt.Printf("Product ID: %d\n", product.Id)
	fmt.Printf("Name: %s\n", product.Name)
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("Price: $%.2f\n", product.Price)
	fmt.Printf("Stock: %d\n", product.Stock)
}

func (product Product) isAvailable() bool {
	if product.Stock > 0 {
		return true
	}
	return false
}
