package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title      string
	price      money
	publishing interface{}
}

func (b book) print() {
	p := format(b.publishing)
	fmt.Printf("%-13s : %s - (%v)\n", b.title, b.price.string(), p)
}

func format(publishing interface{}) string {
	if publishing == nil {
		return "unknown"
	}

	var t int
	if publishing, ok := publishing.(int); ok {
		t = publishing
	}
	if publishing, ok := publishing.(string); ok {
		t, _ = strconv.Atoi(publishing)
	}

	u := time.Unix(int64(t), 0)
	return u.String()

}
