package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			defer group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}
