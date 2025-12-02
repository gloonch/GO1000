package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}

	// semicolon remains because of the hidden true statement
	//switch i := 10; true{
	switch i := 10; {
	case i > 0:
		fmt.Printf("positive")
	case i < 0:
		fmt.Printf("negative")
	default:
		fmt.Printf("zero")
	}

}
