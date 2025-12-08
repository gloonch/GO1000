package main

import "fmt"

// an abstract type/protocol/contract **no implementation
type printer interface {
	print()
}

type list []printer

func (items list) print() {
	if len(items) == 0 {
		fmt.Println("Sorry, nothing to play")

		return
	}

	for _, l := range items {
		//fmt.Printf("(%-10T) --> ", l)
		l.print()
	}
}
