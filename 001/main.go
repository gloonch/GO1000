package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	leng := len(msg)

	exMark := strings.Repeat("!", leng)
	var sliceOfStr []string
	finalStr := append(sliceOfStr, exMark, msg, exMark)

	fmt.Println(finalStr)

}
