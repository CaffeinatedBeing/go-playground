# Performance Optimization Challenge: Finding Duplicates

## Problem Description

You are given a large array of integers. The current implementation for finding duplicates is slow and inefficient, especially for large datasets. Your task is to:

1. Identify the performance bottleneck in the provided code
2. Optimize the algorithm to efficiently find all duplicate numbers
3. Compare the performance before and after optimization

## Learning Objectives

- Profiling Go code for performance
- Algorithmic optimization
- Efficient use of Go data structures
- Benchmarking and measuring improvements

## Getting Started

1. Read the code in `main.go` and run the provided benchmarks
2. Use Go's benchmarking tools to measure performance
3. Optimize the code to improve performance
4. Verify correctness and performance using the provided tests and benchmarks

## Hints

- Consider the time complexity of the current approach
- Explore using maps or sets for faster lookups
- Use Go's built-in benchmarking tools (`go test -bench .`)

## Success Criteria

- The optimized solution should be significantly faster on large datasets
- All tests must pass
- The code should be clean and idiomatic

## Bonus Challenge

- Parallelize the duplicate finding algorithm for even better performance
- Profile memory usage and optimize allocations
