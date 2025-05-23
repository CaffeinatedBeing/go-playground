package main

import "fmt"

// FindDuplicatesOptimized finds duplicates using a map for O(n) performance
func FindDuplicatesOptimized(nums []int) []int {
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			duplicates[n] = true
		} else {
			seen[n] = true
		}
	}
	result := make([]int, 0, len(duplicates))
	for d := range duplicates {
		result = append(result, d)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 8, 5, 9, 10, 2}
	dupes := FindDuplicatesOptimized(nums)
	fmt.Println("Duplicates:", dupes)
}
