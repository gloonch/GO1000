package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var mu sync.Mutex

	for i := 0; i <= 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(count)
			count++
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(count)
}
