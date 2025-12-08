package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	publishing interface{}
}

func (b book) print() {
	p := format(b.publishing)
	fmt.Printf("%-13s : %s - (%v)\n", b.title, b.price.string(), p)
}

func format(publishing interface{}) string {

	var t int

	switch publishing := publishing.(type) {
	case string:
		t, _ = strconv.Atoi(publishing)
	case int:
		t = publishing
	default:
		return "unknown"
	}

	const layout = "2006/01"
	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
