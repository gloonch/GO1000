package main

import (
	"encoding/json"
	"fmt"
)

type Permissions map[string]bool

type user struct {
	Name        string `json:"username"`
	Password    string `json:"-"`
	Permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"imdadkhan", "123456", nil},
		{"god", "777777", Permissions{
			"admin": true,
		}},
		{"devil", "666666", Permissions{
			"write": true,
		}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("error ocurred: ", err)

		return
	}

	fmt.Println(string(out))

}
