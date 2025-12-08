package main

import (
	"fmt"
)

func main() {
	var (
		counter int
		V       int
		P       *int
	)

	counter = 100
	P = &counter
	V = *P

	fmt.Printf("The counter is: %-13d address: %-13p\n", counter, &counter)
	fmt.Printf("P is          : %-13p address: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V is          : %-13d address: %-13p\n", V, &V)

	fmt.Println("\n ---------- changing the counter ----------")

	counter = 10

	fmt.Printf("The counter is: %-13d address: %-13p\n", counter, &counter)

	fmt.Println("\n ---------- changing counter value in func passValue() ----------")
	counter = passValue(counter)
	fmt.Printf("The counter is: %-13d address: %-13p\n", counter, &counter)

	fmt.Println("\n ---------- changing counter value in func passPointerValue() ----------")
	// pointer variables inside the function are local
	// they get new memory addresses each time the function gets called
	passPointerValue(&counter)
	passPointerValue(&counter)
	passPointerValue(&counter)
	fmt.Printf("The counter is: %-13d address: %-13p\n", counter, &counter)

}

func passPointerValue(i *int) {
	fmt.Printf("i is          : %-13p address: %-13p *i: %-13d\n",
		i, &i, *i)

	*i++

	fmt.Printf("i is          : %-13p address: %-13p *i: %-13d\n",
		i, &i, *i)
}

func passValue(n int) int {
	n = 50
	fmt.Printf("n is         : %-13d address: %-13p\n", n, &n)

	return n
}
