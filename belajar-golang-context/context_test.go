package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
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

func CreateCounter(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)
	group.Add(1)
	go func() {
		defer close(destination)
		defer group.Done()
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //SIMULASI SLOW
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	group := &sync.WaitGroup{}
	destination := CreateCounter(ctx, group)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter:", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim signal cancel ke context
	group.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	group := &sync.WaitGroup{}
	destination := CreateCounter(ctx, group)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter:", n)
	}
	group.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()
	group := &sync.WaitGroup{}
	destination := CreateCounter(ctx, group)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter:", n)
	}
	group.Wait()
	fmt.Println(runtime.NumGoroutine())
}
