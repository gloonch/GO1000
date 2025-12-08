package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type DBConnection struct {
	index int32
}

func main() {

	var count int32

	connectionPool := &sync.Pool{
		New: func() any {
			return &DBConnection{
				index: atomic.AddInt32(&count, 1),
			}
		},
	}

	for i := 0; i < 10; i++ {
		go func() {

			con := connectionPool.Get().(*DBConnection)
			defer connectionPool.Put(con)
			fmt.Println("con : %v\n", con)

		}()
	}

	fmt.Scanln()
}
