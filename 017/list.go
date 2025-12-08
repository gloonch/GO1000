package main

import "fmt"

type list []*product

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

func (items list) discount(ratio float64) {
	items.discount(ratio)
}
