package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	max   = 5
	usage = `Welcome to Lucky Numbers!

The program will pick %d random numbers.
Your mission is to guess on of those numbers.

The greater the number is, the harder it gets.

Wanna play?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, max)
		return // stops the program and u have to run it again
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if guess < 0 {
		fmt.Println("Please provide a positive number.")
		return
	}

	for turn := 0; turn <= max; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("YOU WON!")
			return
		}
	}
	fmt.Println("YOU LOST... Try again?")
}
