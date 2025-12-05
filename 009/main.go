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

	var names []byte

	for _, file := range files {
		if info, err := file.Info(); err == nil && info.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		} else if err != nil {
			fmt.Println("Error ocurred while getting file info: \n", err)
		}
	}
	if err := os.WriteFile("empty_files_list.txt", names, 0644); err != nil {
		fmt.Println("Error ocurred while writing to file: \n", err)

		return
	} else {
		fmt.Println("Empty files list written to file successfully")
	}

	fmt.Println(string(names))
}
