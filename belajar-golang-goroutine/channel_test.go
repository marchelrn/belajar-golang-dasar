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

func Response(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Marchel Matthew Manullang"
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
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Data ke-%d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string, 2)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go Response(channel1)
	go Response(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestSelectChannelDefault(t *testing.T) {
	channel1 := make(chan string, 2)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go Response(channel1)
	go Response(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}
		if counter == 2 {
			break
		}
	}
}
