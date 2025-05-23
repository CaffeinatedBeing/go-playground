package main

import (
	"testing"
)

func BenchmarkDataProcessing(b *testing.B) {
	// Generate test data once
	data := generateTestData(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processor := NewDataProcessor()
		for _, point := range data {
			processor.ProcessDataPoint(point)
		}
		_ = processor.GetStatistics()
	}
}

func BenchmarkTagProcessing(b *testing.B) {
	// Create a processor with some initial data
	processor := NewDataProcessor()
	initialData := generateTestData(1000)
	for _, point := range initialData {
		processor.ProcessDataPoint(point)
	}

	// Generate new test data for each iteration
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newPoint := DataPoint{
			ID:    i,
			Value: float64(i),
			Tags:  []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
		}
		processor.ProcessDataPoint(newPoint)
	}
}

func BenchmarkStatisticsGeneration(b *testing.B) {
	// Create a processor with some data
	processor := NewDataProcessor()
	data := generateTestData(10000)
	for _, point := range data {
		processor.ProcessDataPoint(point)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processor.GetStatistics()
	}
}

func BenchmarkDataGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateTestData(1000)
	}
}

func TestDataProcessorCorrectness(t *testing.T) {
	processor := NewDataProcessor()

	// Test data with known tags
	testData := []DataPoint{
		{ID: 1, Value: 10.0, Tags: []string{"tag1", "tag2"}},
		{ID: 2, Value: 20.0, Tags: []string{"tag2", "tag3"}},
		{ID: 3, Value: 30.0, Tags: []string{"tag1", "tag3"}},
	}

	// Process the test data
	for _, point := range testData {
		processor.ProcessDataPoint(point)
	}

	// Get statistics
	stats := processor.GetStatistics()

	// Verify results
	if stats["total_points"] != 3 {
		t.Errorf("Expected 3 total points, got %d", stats["total_points"])
	}

	if stats["unique_tags"] != 3 {
		t.Errorf("Expected 3 unique tags, got %d", stats["unique_tags"])
	}

	tagCounts := stats["tag_counts"].(map[string]int)
	expectedCounts := map[string]int{
		"tag1": 2,
		"tag2": 2,
		"tag3": 2,
	}

	for tag, expected := range expectedCounts {
		if count := tagCounts[tag]; count != expected {
			t.Errorf("Expected count %d for tag %s, got %d", expected, tag, count)
		}
	}
}
