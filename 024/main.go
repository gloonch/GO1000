package main

import "fmt"

func main() {

	ch := make(chan string)

	// receiver
	go func() {
		message := <-ch
		fmt.Println(message)
	}()

	// sender
	go func() {
		ch <- "Hello World!"
	}()

	fmt.Scanln()

}
