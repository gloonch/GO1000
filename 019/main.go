package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex   sync.Mutex
	cond    = sync.NewCond(&mutex)
)

func main() {
	go producer()
	go consumer()

	time.Sleep(5 * time.Second)
}

func producer() {
	for {
		mutex.Lock()
		if counter > 0 {
			cond.Wait()
		}
		time.Sleep(1 * time.Second)
		counter++
		fmt.Println("incremented: ", counter)
		mutex.Unlock()
		cond.Signal()
	}

}

func consumer() {
	for {
		mutex.Lock()
		if counter == 0 {
			cond.Wait()
		}
		time.Sleep(1 * time.Second)
		counter--
		fmt.Println("decreased: ", counter)
		mutex.Unlock()
		cond.Signal()
	}
}
