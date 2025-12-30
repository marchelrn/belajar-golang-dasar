package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New Data"
		},
	}
	group := sync.WaitGroup{}
	pool.Put("Marchel")
	pool.Put("Matthew")
	pool.Put("Manullang")
	for i := 0; i <= 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			if data != nil {
				fmt.Println(data)
				time.Sleep(1 * time.Second)
				pool.Put(data)
			}
		}()
	}
	group.Wait()
	fmt.Println("Selesai")
}
