package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(wait *sync.WaitGroup) {
	defer wait.Done()

	wait.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	wait := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsync(wait)
	}
	wait.Wait()
	fmt.Println("Selesai")
}

func RunTask(Name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Menjalankan Task:", Name)
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Menjalankan Task:", Name)
}

func TestRunTask(t *testing.T) {
	wg := &sync.WaitGroup{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
		"Task 4",
		"Task 5",
	}
	for _, task := range tasks {
		wg.Add(1)
		go RunTask(task, wg)
	}
	fmt.Println("Sistem : Menunggu Semua Task Selesai")
	wg.Wait()
	fmt.Println("Sistem : Semua Task Selesai")
}
