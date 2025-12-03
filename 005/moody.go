package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[enter your name]")
		return
	}

	name := args[0]

	moods := [...]string{
		"happy", "sad", "awesome", "bad", "terrible", "good",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[n])
}
