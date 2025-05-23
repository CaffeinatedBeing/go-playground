# Slice Capacity and Memory Allocation Challenge

## Problem Description

You are given a simple data processing pipeline that processes large datasets using slices. The current implementation has performance issues due to inefficient memory allocation and slice capacity management. Your task is to:

1. Identify where memory is being allocated inefficiently
2. Optimize slice capacity management
3. Reduce memory allocations
4. Improve overall performance

## Learning Objectives

- Understanding Go's slice internals
- Memory allocation patterns
- Slice capacity growth
- Memory profiling
- Performance optimization

## Getting Started

1. Read through the code in `main.go`
2. Run the program and observe the memory usage
3. Use `go tool pprof` to identify memory allocation hotspots
4. Implement your solution
5. Verify the improvements using the provided benchmarks

## Hints

1. Look at how slices are being created and grown
2. Consider pre-allocating slices with appropriate capacity
3. Think about when to use `make` with capacity
4. Consider the trade-offs between memory usage and performance

## Success Criteria

- Reduced memory allocations (verify using pprof)
- Improved performance (verify using benchmarks)
- Maintained correctness of the data processing
- Clear documentation of optimization decisions

## Bonus Challenge

1. Implement a custom slice type with optimized growth strategy
2. Add memory usage monitoring
3. Create a memory-efficient iterator pattern
4. Implement parallel processing with controlled memory usage
