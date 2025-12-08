package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// var count atomic.Int32
	var count int32
	var wg sync.WaitGroup

	iterations := 10000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			atomic.AddInt32(&count, 1)
			// count.Add(1)

		}()
	}

	wg.Wait()

	// fmt.Println(count.Load())
	fmt.Println(atomic.LoadInt32(&count))
}
