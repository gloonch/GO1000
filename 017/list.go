package main

import "fmt"

// an abstract type/protocol/contract **no implementation
type item interface {
	print()
	discount(ratio float64)
}

type list []item

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
	type discounter interface {
		discount(ratio float64)
	}

	for _, l := range items {
		//if l, hasDiscount := l.(interface {discount(ratio float64)}); hasDiscount { // using interface inline
		//g, isGame := l.(*game) { // finding out the dynamic value of an interface
		if l, hasDiscount := l.(discounter); hasDiscount {
			l.discount(ratio)
		}
	}
}
