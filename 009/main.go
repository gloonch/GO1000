package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Provide a directory path")

		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println("Error ocurred while reading directory: \n", err)

		return
	}

	for _, file := range files {
		if info, err := file.Info(); err == nil && info.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		} else if err != nil {
			fmt.Println("Error ocurred while getting file info: \n", err)
		}
	}
}
