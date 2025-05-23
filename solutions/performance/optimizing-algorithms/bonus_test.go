package main

import (
	"math/rand"
	"testing"
)

func TestFindDuplicatesParallel(t *testing.T) {
	nums := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 8, 5, 9, 10, 2}
	dupes := FindDuplicatesParallel(nums)
	expected := map[int]bool{1: true, 2: true, 5: true}
	if len(dupes) != len(expected) {
		t.Errorf("Expected %d duplicates, got %d", len(expected), len(dupes))
	}
	for _, d := range dupes {
		if !expected[d] {
			t.Errorf("Unexpected duplicate: %d", d)
		}
	}
}

func generateLargeSliceParallel(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(size / 2) // force some duplicates
	}
	return nums
}

func BenchmarkFindDuplicatesParallel(b *testing.B) {
	nums := generateLargeSliceParallel(2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindDuplicatesParallel(nums)
	}
}
