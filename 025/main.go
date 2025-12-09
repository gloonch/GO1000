package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go send(ch)
	go send(ch)
	go send(ch)
	go receive(ch)
	go receive(ch)
	go receive(ch)

	fmt.Scanln()
}

func send(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("hello world number %d", i)
	}
}

func receive(ch <-chan string) {
	for msg := range ch {
		time.Sleep(time.Second)
		fmt.Println(msg)
	}
}
