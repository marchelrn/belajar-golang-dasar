package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("background:", background)

	todo := context.TODO()
	fmt.Println("todo:", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println("contextA:", contextA)
	fmt.Println("contextB:", contextB)
	fmt.Println("contextC:", contextC)
	fmt.Println("contextD:", contextD)
	fmt.Println("contextE:", contextE)
	fmt.Println("contextF:", contextF)
	fmt.Println("contextG:", contextG)

	fmt.Println(contextF.Value("f")) // output: F
	fmt.Println(contextF.Value("c")) // output: C
	fmt.Println(contextF.Value("b")) // output: <nil>, karena contextF berasal dari contextC, bukan contextB
	fmt.Println(contextF.Value("b")) // output: <nil>, karena method value() tidak bisa mengambil value dari childnya
}
