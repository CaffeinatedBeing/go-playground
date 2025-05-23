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
	// BUG: No capacity specified for these slices
	processedData []DataPoint
	uniqueTags    []string
	tagCounts     map[string]int
}

// NewDataProcessor creates a new data processor
func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		// BUG: No initial capacity specified
		processedData: []DataPoint{},
		uniqueTags:    []string{},
		tagCounts:     make(map[string]int),
	}
}

// ProcessDataPoint processes a single data point
func (p *DataProcessor) ProcessDataPoint(point DataPoint) {
	// BUG: Inefficient slice growth
	p.processedData = append(p.processedData, point)

	// BUG: Inefficient tag processing
	for _, tag := range point.Tags {
		p.tagCounts[tag]++
		// BUG: Inefficient unique tag collection
		if !contains(p.uniqueTags, tag) {
			p.uniqueTags = append(p.uniqueTags, tag)
		}
	}
}

// GetStatistics returns processing statistics
func (p *DataProcessor) GetStatistics() map[string]interface{} {
	// BUG: Inefficient map creation
	stats := make(map[string]interface{})

	// BUG: Inefficient slice operations
	stats["total_points"] = len(p.processedData)
	stats["unique_tags"] = len(p.uniqueTags)
	stats["tag_counts"] = p.tagCounts

	return stats
}

// contains checks if a string is in a slice
func contains(slice []string, str string) bool {
	// BUG: Inefficient linear search
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// generateTestData creates a test dataset
func generateTestData(size int) []DataPoint {
	// BUG: No capacity specified for the result slice
	var data []DataPoint

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		// BUG: Inefficient tag generation
		tags := []string{}
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
