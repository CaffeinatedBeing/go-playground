package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DataPoint represents a single data point in our dataset
type DataPoint struct {
	ID    int
	Value float64
	Tags  []string
}

// DataProcessor processes a stream of data points
type DataProcessor struct {
	processedData []DataPoint
	uniqueTags    map[string]struct{} // Using map for O(1) lookups
	tagCounts     map[string]int
}

// NewDataProcessor creates a new data processor
func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		processedData: make([]DataPoint, 0, 1000), // Pre-allocate with reasonable capacity
		uniqueTags:    make(map[string]struct{}),
		tagCounts:     make(map[string]int),
	}
}

// ProcessDataPoint processes a single data point
func (p *DataProcessor) ProcessDataPoint(point DataPoint) {
	// Pre-allocate tags slice with known capacity
	tags := make([]string, 0, len(point.Tags))

	// Process tags efficiently
	for _, tag := range point.Tags {
		// Update tag counts
		p.tagCounts[tag]++

		// Add to unique tags if not present
		if _, exists := p.uniqueTags[tag]; !exists {
			p.uniqueTags[tag] = struct{}{}
		}

		tags = append(tags, tag)
	}

	// Update point's tags with pre-allocated slice
	point.Tags = tags

	// Append to processed data
	p.processedData = append(p.processedData, point)
}

// GetStatistics returns processing statistics
func (p *DataProcessor) GetStatistics() map[string]interface{} {
	// Pre-allocate map with known size
	stats := make(map[string]interface{}, 3)

	stats["total_points"] = len(p.processedData)
	stats["unique_tags"] = len(p.uniqueTags)
	stats["tag_counts"] = p.tagCounts

	return stats
}

// generateTestData creates a test dataset
func generateTestData(size int) []DataPoint {
	// Pre-allocate slice with known capacity
	data := make([]DataPoint, 0, size)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		// Pre-allocate tags slice with maximum possible size
		tags := make([]string, 0, 5)
		numTags := rand.Intn(5) + 1

		for j := 0; j < numTags; j++ {
			tags = append(tags, fmt.Sprintf("tag%d", rand.Intn(10)))
		}

		data = append(data, DataPoint{
			ID:    i,
			Value: rand.Float64() * 100,
			Tags:  tags,
		})
	}

	return data
}

func main() {
	// Generate a large dataset
	data := generateTestData(100000)

	// Create and use the processor
	processor := NewDataProcessor()

	start := time.Now()

	// Process all data points
	for _, point := range data {
		processor.ProcessDataPoint(point)
	}

	// Get and print statistics
	stats := processor.GetStatistics()

	duration := time.Since(start)

	fmt.Printf("Processing completed in %v\n", duration)
	fmt.Printf("Total points processed: %d\n", stats["total_points"])
	fmt.Printf("Unique tags found: %d\n", stats["unique_tags"])
	fmt.Printf("Tag counts: %v\n", stats["tag_counts"])
}
