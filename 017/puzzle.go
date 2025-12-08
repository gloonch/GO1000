package main

import "fmt"

type puzzle struct {
	title string
	price money
}

func (p puzzle) print() {
	fmt.Printf("%-13s : %s\n", p.title, p.price.string())
}
