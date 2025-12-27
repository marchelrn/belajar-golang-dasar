package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

// goroutine dengan parameter, kode ini menunujukan bahwa goroutine sangat ringan
func displayNumber(number int) {
	fmt.Println("Number:", number)
}

func TestDisplay(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go displayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
