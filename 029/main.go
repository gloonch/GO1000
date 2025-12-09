package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var success, failure chan string

func main() {

	success = make(chan string)
	failure = make(chan string)

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

	go func() {
		for {
			select {
			case success := <-success:
				fmt.Println(success)
				wg.Done()
			case failure := <-failure:
				fmt.Println(failure)
				wg.Done()

			}
		}
	}()

	wg.Wait()
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		failure <- fmt.Sprintf("Error fetching %s: %s\n", url, err.Error())

		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		failure <- fmt.Sprintf("non-success on %s: %v\n", url, res.StatusCode)

		return
	}

	success <- fmt.Sprintf("success on %v : %v\n", url, res.StatusCode)
}
