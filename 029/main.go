package main

import (
	"fmt"
	"net/http"
	"sync"
)

type response struct {
	url     string
	code    int
	message string
}

var wg sync.WaitGroup
var success, failure chan response

func main() {

	success = make(chan response)
	failure = make(chan response)

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
		failure <- response{
			url:     url,
			code:    0,
			message: err.Error(),
		}

		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		failure <- response{
			url:     url,
			code:    res.StatusCode,
			message: res.Status,
		}

		return
	}

	success <- response{
		url:     url,
		code:    res.StatusCode,
		message: res.Status,
	}
}
