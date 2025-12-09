package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
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

type appStatistics struct {
	successResponses atomic.Int64
	failureResponses atomic.Int64
	idleTime         atomic.Int64
	requestTime      atomic.Int64
	appTotalTime     atomic.Int64
}

var wg sync.WaitGroup
var mu sync.Mutex
var success, failure chan response
var sumResponse []*summarizedResponse
var statistics appStatistics

func main() {

	start := time.Now()

	success = make(chan response)
	failure = make(chan response)

	statistics = appStatistics{}

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
			startTime := time.Now()
			select {
			case success := <-success:
				summarize(success)
				statistics.successResponses.Add(1)
				fmt.Println(success)
				wg.Done()
			case failure := <-failure:
				summarize(failure)
				statistics.failureResponses.Add(1)
				fmt.Println(failure)
				wg.Done()
			default:
				// idle time
				idleTime := time.Since(startTime).Microseconds()
				statistics.idleTime.Add(idleTime)

			}
		}
	}()

	wg.Wait()

	statistics.appTotalTime.Add(time.Since(start).Microseconds())

	fmt.Println("\n--- results ---")
	for _, sum := range sumResponse {
		fmt.Printf("results by code: %v\n", sum.code)
	}

	fmt.Println("\n--- statistics ---")
	fmt.Printf("successes: %v\n", statistics.successResponses.Load())
	fmt.Printf("failures: %v\n", statistics.failureResponses.Load())
	fmt.Printf("idle time (sec): %v\n", float64(statistics.idleTime.Load())/1000/1000)
	fmt.Printf("request time (sec): %v\n", float64(statistics.requestTime.Load())/1000/1000)
	fmt.Printf("app total time (sec): %v\n", float64(statistics.appTotalTime.Load())/1000/1000)

}

func fetch(url string) {
	start := time.Now()
	defer func() {
		requestTime := time.Since(start).Microseconds()
		statistics.requestTime.Add(requestTime)
	}()
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
