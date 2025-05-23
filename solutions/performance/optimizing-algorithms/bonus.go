package main

import (
	"fmt"
	"runtime"
	"sync"
)

// FindDuplicatesParallel finds duplicates using goroutines for parallelism
func FindDuplicatesParallel(nums []int) []int {
	workers := runtime.NumCPU()
	chunkSize := (len(nums) + workers - 1) / workers

	seenCh := make(chan map[int]bool, workers)
	dupCh := make(chan map[int]bool, workers)
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		start := w * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			seen := make(map[int]bool)
			dups := make(map[int]bool)
			for _, n := range chunk {
				if seen[n] {
					dups[n] = true
				} else {
					seen[n] = true
				}
			}
			seenCh <- seen
			dupCh <- dups
		}(nums[start:end])
	}

	wg.Wait()
	close(seenCh)
	close(dupCh)

	// Merge results
	globalSeen := make(map[int]bool)
	globalDups := make(map[int]bool)
	for seen := range seenCh {
		for n := range seen {
			if globalSeen[n] {
				globalDups[n] = true
			} else {
				globalSeen[n] = true
			}
		}
	}
	for dups := range dupCh {
		for n := range dups {
			globalDups[n] = true
		}
	}

	result := make([]int, 0, len(globalDups))
	for d := range globalDups {
		result = append(result, d)
	}
	return result
}

func mainParallel() {
	nums := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 8, 5, 9, 10, 2}
	dupes := FindDuplicatesParallel(nums)
	fmt.Println("Duplicates (parallel):", dupes)
}

// ---
// Memory Profiling & Allocation Optimization Notes:
//
// - Use `go test -bench . -benchmem` to measure allocations.
// - Use `go test -cpuprofile cpu.prof -memprofile mem.prof` and `go tool pprof` for profiling.
// - Pre-allocate slices/maps when possible to reduce allocations.
// - Avoid unnecessary intermediate slices in hot code paths.
// - For very large datasets, consider streaming/iterator patterns to reduce memory footprint.
