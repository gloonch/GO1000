package main

import "fmt"

type list []*game

func (items list) print() {
	if len(items) == 0 {
		fmt.Println("Sorry, nothing to play")

		return
	}

	for _, l := range items {
		l.print()
	}
}
