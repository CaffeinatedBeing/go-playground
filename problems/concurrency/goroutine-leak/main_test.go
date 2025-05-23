package main

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestCrawlerNoLeaks(t *testing.T) {
	// Record initial number of goroutines
	initialGoroutines := runtime.NumGoroutine()

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://invalid-url-that-will-fail.com",
		"https://google.com",
	}

	// Create a context with a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	crawler := NewCrawler(urls)
	crawler.Start(ctx)

	// Collect all results
	var results []Result
	for result := range crawler.Results() {
		results = append(results, result)
	}

	// Give some time for goroutines to clean up
	time.Sleep(100 * time.Millisecond)

	// Check final number of goroutines
	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines > initialGoroutines {
		t.Errorf("Goroutine leak detected! Initial: %d, Final: %d",
			initialGoroutines, finalGoroutines)
	}

	// Verify we got results for all URLs
	if len(results) != len(urls) {
		t.Errorf("Expected %d results, got %d", len(urls), len(results))
	}
}

func TestCrawlerContextCancellation(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://google.com",
	}

	// Create a context that we'll cancel immediately
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	crawler := NewCrawler(urls)
	crawler.Start(ctx)

	// Collect results (should be empty or partial due to cancellation)
	var results []Result
	for result := range crawler.Results() {
		results = append(results, result)
	}

	// Give some time for goroutines to clean up
	time.Sleep(100 * time.Millisecond)

	// Check final number of goroutines
	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines > initialGoroutines {
		t.Errorf("Goroutine leak detected after cancellation! Initial: %d, Final: %d",
			initialGoroutines, finalGoroutines)
	}
}

func TestCrawlerTimeout(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()

	// Add a very slow URL that will timeout
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://httpstat.us/200?sleep=5000", // This will sleep for 5 seconds
	}

	// Create a context with a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	crawler := NewCrawler(urls)
	crawler.Start(ctx)

	// Collect results
	var results []Result
	for result := range crawler.Results() {
		results = append(results, result)
	}

	// Give some time for goroutines to clean up
	time.Sleep(100 * time.Millisecond)

	// Check final number of goroutines
	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines > initialGoroutines {
		t.Errorf("Goroutine leak detected after timeout! Initial: %d, Final: %d",
			initialGoroutines, finalGoroutines)
	}
}
