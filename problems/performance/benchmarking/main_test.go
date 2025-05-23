package main

import (
	"fmt"
	"runtime"
	"testing"
)

// # Run all benchmarks with memory stats
// go test -bench=. -benchmem

// # Run a specific benchmark
// go test -bench=BenchmarkStringProcessors -benchmem

// # Run with longer duration
// go test -bench=. -benchmem -benchtime=5s

// # Run with CPU profiling
// go test -bench=. -cpuprofile=cpu.prof

// # Run with memory profiling
// go test -bench=. -memprofile=mem.prof

// ExampleBenchmarkCommands demonstrates how to run the benchmarks
func ExampleBenchmarkCommands() {
	fmt.Println("To run all benchmarks:")
	fmt.Println("go test -bench=. -benchmem")
	fmt.Println("\nTo run a specific benchmark:")
	fmt.Println("go test -bench=BenchmarkStringProcessors -benchmem")
	fmt.Println("\nTo run benchmarks with longer duration:")
	fmt.Println("go test -bench=. -benchmem -benchtime=5s")
	fmt.Println("\nTo run benchmarks with CPU profile:")
	fmt.Println("go test -bench=. -cpuprofile=cpu.prof")
	fmt.Println("\nTo run benchmarks with memory profile:")
	fmt.Println("go test -bench=. -memprofile=mem.prof")
}

func TestStringProcessors(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"single char", "a", "A"},
		{"single char upper", "A", "a"},
		{"mixed case", "Hello", "hELLO"},
		{"all upper", "HELLO", "hello"},
		{"all lower", "hello", "HELLO"},
	}

	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, p := range processors {
		t.Run(p.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := p.processor.Process(tc.input)
					if result != tc.expected {
						t.Errorf("Expected %q, got %q", tc.expected, result)
					}
				})
			}
		})
	}
}

// BenchmarkStringProcessors benchmarks different string processors with various input sizes
func BenchmarkStringProcessors(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, size := range sizes {
		input := GenerateTestString(size)
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for _, p := range processors {
				b.Run(p.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = p.processor.Process(input)
					}
				})
			}
		})
	}
}

// BenchmarkStringProcessorsAllocs benchmarks memory allocations for different processors
func BenchmarkStringProcessorsAllocs(b *testing.B) {
	input := GenerateTestString(1000)
	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, p := range processors {
		b.Run(p.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = p.processor.Process(input)
			}
		})
	}
}

// BenchmarkStringProcessorsParallel benchmarks parallel processing performance
func BenchmarkStringProcessorsParallel(b *testing.B) {
	input := GenerateTestString(1000)
	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, p := range processors {
		b.Run(p.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = p.processor.Process(input)
				}
			})
		})
	}
}

// BenchmarkStringProcessorsMetrics benchmarks with custom memory metrics
func BenchmarkStringProcessorsMetrics(b *testing.B) {
	input := GenerateTestString(1000)
	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, p := range processors {
		b.Run(p.name, func(b *testing.B) {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			allocBefore := m.TotalAlloc

			for i := 0; i < b.N; i++ {
				_ = p.processor.Process(input)
			}

			runtime.ReadMemStats(&m)
			allocAfter := m.TotalAlloc
			bytesPerOp := float64(allocAfter-allocBefore) / float64(b.N)
			b.ReportMetric(bytesPerOp, "B/op")
		})
	}
}

// BenchmarkStringProcessorsScenarios benchmarks different input scenarios
func BenchmarkStringProcessorsScenarios(b *testing.B) {
	scenarios := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"SingleChar", "a"},
		{"MixedCase", "HelloWorld"},
		{"AllUpper", "HELLOWORLD"},
		{"AllLower", "helloworld"},
		{"LongString", GenerateTestString(1000)},
	}

	processors := []struct {
		name      string
		processor StringProcessor
	}{
		{"Naive", &NaiveProcessor{}},
		{"Builder", &BuilderProcessor{}},
		{"Bytes", &BytesProcessor{}},
		{"Prealloc", &PreallocProcessor{}},
		{"Map", NewMapProcessor()},
	}

	for _, s := range scenarios {
		b.Run(s.name, func(b *testing.B) {
			for _, p := range processors {
				b.Run(p.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = p.processor.Process(s.input)
					}
				})
			}
		})
	}
}
