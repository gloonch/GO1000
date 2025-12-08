package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() {
		return "unknown"
	}
	return ts.Format("2006-01-02")
}

func toTimestamp(publishing interface{}) (ts timestamp) {

	var t int

	switch publishing := publishing.(type) {
	case string:
		t, _ = strconv.Atoi(publishing)
	case int:
		t = publishing

	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
