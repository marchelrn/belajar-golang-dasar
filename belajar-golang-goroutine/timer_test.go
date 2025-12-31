package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Time Now :", time.Now())

	time := <-timer.C
	fmt.Println("Time :", time)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)
	fmt.Println("Time Now :", time.Now())

	time := <-timer
	fmt.Println("Time :", time)
}
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("After Func :", time.Now())
		group.Done()
	})
	fmt.Println("After Func :", time.Now())
	group.Wait()
}
