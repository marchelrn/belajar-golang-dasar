package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time)
		if time.Second()%5 == 0 {
			ticker.Stop()
			break
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
		if time.Second()%5 == 0 {
			break
		}
	}
}

func TestPolling(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := time.After(1 * time.Second)
	fmt.Println("Mulai Polling:", time.Now())

	done := true
	for done {
		select {
		case t := <-ticker.C:
			fmt.Println("Status Server OK:", t.Format("15:04:05"))
		case <-timeout:
			fmt.Println("Polling Selesai:", time.Now())
			done = false
		}
	}
	fmt.Println("Aplikasi selesai")
}
