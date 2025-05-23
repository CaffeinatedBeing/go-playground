package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result represents the result of a URL fetch
type Result struct {
	URL     string
	Status  int
	Error   error
	Latency time.Duration
}

// fetchURL fetches a single URL and returns the result
func fetchURL(url string) Result {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return Result{
			URL:     url,
			Error:   err,
			Latency: time.Since(start),
		}
	}
	defer resp.Body.Close()

	return Result{
		URL:     url,
		Status:  resp.StatusCode,
		Latency: time.Since(start),
	}
}

// Crawler is a simple web crawler that fetches URLs concurrently
type Crawler struct {
	urls    []string
	results chan Result
	wg      sync.WaitGroup
}

// NewCrawler creates a new crawler instance
func NewCrawler(urls []string) *Crawler {
	return &Crawler{
		urls:    urls,
		results: make(chan Result),
	}
}

// Start begins the crawling process
func (c *Crawler) Start(ctx context.Context) {
	for _, url := range c.urls {
		c.wg.Add(1)
		go func(url string) {
			// BUG: This goroutine might leak if the context is cancelled
			// before the fetch completes
			result := fetchURL(url)
			c.results <- result
			c.wg.Done()
		}(url)
	}

	// BUG: This goroutine might leak if the context is cancelled
	go func() {
		c.wg.Wait()
		close(c.results)
	}()
}

// Results returns the channel of results
func (c *Crawler) Results() <-chan Result {
	return c.results
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://invalid-url-that-will-fail.com",
		"https://google.com",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	crawler := NewCrawler(urls)
	crawler.Start(ctx)

	// BUG: This loop might not exit if the context is cancelled
	// because the results channel might never be closed
	for result := range crawler.Results() {
		if result.Error != nil {
			fmt.Printf("Error fetching %s: %v (took %v)\n",
				result.URL, result.Error, result.Latency)
			continue
		}
		fmt.Printf("Successfully fetched %s: %d (took %v)\n",
			result.URL, result.Status, result.Latency)
	}

	// BUG: We're not properly waiting for all goroutines to finish
	// before the program exits
	fmt.Println("Crawling completed!")
}
