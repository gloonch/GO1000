package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	channel4 := make(chan string)
	channel5 := make(chan string)

	go func() {
		channel1 <- "message to channel 1"
	}()

	go func() {
		channel2 <- "message to channel 2"
	}()

	go func() {
		channel3 <- "message to channel 3"
	}()

	go func() {
		channel4 <- "message to channel 4"
	}()

	go func() {
		channel5 <- "message to channel 5"
	}()

	timeout := time.After(3 * time.Second)

	for {
		select {
		case message := <-channel1:
			fmt.Println("message received on channel 1 : ", message)
		case message := <-channel2:
			fmt.Println("message received on channel 2 : ", message)
		case message := <-channel3:
			fmt.Println("message received on channel 3 : ", message)
		case message := <-channel4:
			fmt.Println("message received on channel 4 : ", message)
		case message := <-channel5:
			fmt.Println("message received on channel 5 : ", message)
		case <-timeout:
			fmt.Println("five seconds without getting a new message, going to panic...")
			panic("no message received")
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("no message received")
		}
	}

	fmt.Scanln()
}
