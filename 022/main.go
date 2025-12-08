package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Person struct {
	Name string
	Age  int32
}

func main() {

	var wg sync.WaitGroup
	var person atomic.Value

	person.Store(&Person{Name: "Flashitia", Age: 40}) // it accepts only pointer

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := person.Load().(*Person) // must be casted
			atomic.AddInt32(&p.Age, 1)
		}()
	}

	wg.Wait()

	fmt.Println(person.Load().(*Person))
}
