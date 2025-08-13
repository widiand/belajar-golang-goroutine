package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "eko"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}