package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://google.com/not-found",
		"https://google.com",
		"https://openai.com",
		"https://non-existent-url.com",
	}

	wg.Add(len(urls))
	for _, url := range urls {
		go fetch(url)
	}

	wg.Wait()
}

func fetch(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err.Error())

		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("non-success on %s: %v\n", url, res.StatusCode)

		return
	}

	fmt.Printf("success on %v : %v\n", url, res.StatusCode)
}
