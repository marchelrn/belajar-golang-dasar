package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Marchel Matthew Manullang"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	close(channel)
}

func Response(chann chan string) {
	time.Sleep(2 * time.Second)
	chann <- "Marchel Matthew Manullang"
}

func TestResponse(t *testing.T) {
	channel := make(chan string)
	go Response(channel)
	fmt.Println("Eksekusi")

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	close(channel)
}

func ChannelIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Marchel Matthew Manullang"
}

func ChannelOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestCChannelInOut(t *testing.T) {
	channel := make(chan string)
	go ChannelIn(channel)
	go ChannelOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Marchel"
	channel <- "Matthew"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}
