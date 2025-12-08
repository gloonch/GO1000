package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	iterations := 10000000

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				mutex.Lock()
				count++
				mutex.Unlock()
			}()
		}
	}()

	wg.Wait()

	fmt.Println(count)
}
