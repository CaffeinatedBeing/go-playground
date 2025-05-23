package main

import "fmt"

// FindDuplicatesNaive finds duplicates using a slow O(n^2) approach
func FindDuplicatesNaive(nums []int) []int {
	duplicates := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				// Check if already added
				alreadyAdded := false
				for _, d := range duplicates {
					if d == nums[i] {
						alreadyAdded = true
						break
					}
				}
				if !alreadyAdded {
					duplicates = append(duplicates, nums[i])
				}
			}
		}
	}
	return duplicates
}

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 8, 5, 9, 10, 2}
	dupes := FindDuplicatesNaive(nums)
	fmt.Println("Duplicates:", dupes)
}
