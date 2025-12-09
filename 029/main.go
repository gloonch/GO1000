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

type summarizedResponse struct {
	code        int
	occurrences int
}

var wg sync.WaitGroup
var mu sync.Mutex
var success, failure chan response
var sumResponse []*summarizedResponse

func main() {

	success = make(chan response)
	failure = make(chan response)

	urls := []string{
		"https://google.com/not-found",
		"https://google.com",
		"https://openai.com",
		"https://non-existent-url.com",
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
				summarize(success)
				fmt.Println(success)
				wg.Done()
			case failure := <-failure:
				summarize(failure)
				fmt.Println(failure)
				wg.Done()

			}
		}
	}()

	wg.Wait()

	fmt.Println("--- results ---")
	for _, sum := range sumResponse {
		fmt.Printf("results by code: %v\n", sum.code)
	}
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

func summarize(response response) {
	mu.Lock()
	defer mu.Unlock()

	for _, res := range sumResponse {
		if res.code == response.code {
			res.occurrences++

			return
		}
	}

	sumResponse = append(sumResponse, &summarizedResponse{
		code:        response.code,
		occurrences: 1,
	})

}
