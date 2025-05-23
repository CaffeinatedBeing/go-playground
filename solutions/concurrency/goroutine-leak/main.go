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
func fetchURL(ctx context.Context, url string) Result {
	start := time.Now()

	// Create a new request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Result{
			URL:     url,
			Error:   err,
			Latency: time.Since(start),
		}
	}

	// Use http.DefaultClient with context
	resp, err := http.DefaultClient.Do(req)
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
		results: make(chan Result, len(urls)), // Buffered channel to prevent blocking
	}
}

// Start begins the crawling process
func (c *Crawler) Start(ctx context.Context) {
	for _, url := range c.urls {
		c.wg.Add(1)
		go func(url string) {
			defer c.wg.Done()

			// Use select to handle both context cancellation and fetch completion
			select {
			case <-ctx.Done():
				// Context was cancelled, send error result
				c.results <- Result{
					URL:     url,
					Error:   ctx.Err(),
					Latency: 0,
				}
			default:
				// Fetch the URL
				result := fetchURL(ctx, url)

				// Try to send result, but respect context cancellation
				select {
				case c.results <- result:
				case <-ctx.Done():
					// Context was cancelled while trying to send result
					c.results <- Result{
						URL:     url,
						Error:   ctx.Err(),
						Latency: result.Latency,
					}
				}
			}
		}(url)
	}

	// Close results channel when all goroutines are done
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

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	crawler := NewCrawler(urls)
	crawler.Start(ctx)

	// Process results until channel is closed or context is done
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Crawling cancelled:", ctx.Err())
			return
		case result, ok := <-crawler.Results():
			if !ok {
				// Channel closed, all results processed
				fmt.Println("Crawling completed!")
				return
			}
			if result.Error != nil {
				fmt.Printf("Error fetching %s: %v (took %v)\n",
					result.URL, result.Error, result.Latency)
				continue
			}
			fmt.Printf("Successfully fetched %s: %d (took %v)\n",
				result.URL, result.Status, result.Latency)
		}
	}
}
