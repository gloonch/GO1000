package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("lazy cat jumps again and again and again")

	for k, v := range words {
		fmt.Printf("index %d: %s\n", k, v)
	}
}
