package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if num := os.Args; len(num) != 2 {
		// only has num variable
		fmt.Println("give me a number.")
	} else if n, err := strconv.Atoi(num[1]); err != nil {
		// only has n, num and err variable
		fmt.Printf("Cannot convert %q.\n", num[1])
	} else {
		// has all variables in the if statement
		fmt.Printf("%s * 2 is %d\n", num[1], n*2)
	}

}
