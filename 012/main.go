package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		// input = append(input, '\n')
	}

	var users []user
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println("error ocurred while unmarshalling json object: ", err)

		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name, "\t")

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}
		fmt.Println()
	}
}
