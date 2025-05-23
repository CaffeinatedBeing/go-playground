# Goroutine Leak Challenge

## Problem Description

You are given a simple web crawler implementation that fetches URLs concurrently. However, there's a subtle bug in the code that causes goroutine leaks. Your task is to:

1. Identify where the goroutine leak is occurring
2. Fix the leak while maintaining the concurrent behavior
3. Ensure proper cleanup of resources
4. Add appropriate error handling

## Learning Objectives

- Understanding goroutine lifecycle management
- Proper use of context for cancellation
- Channel cleanup and best practices
- Error handling in concurrent operations
- Debugging goroutine leaks

## Getting Started

1. Read through the code in `main.go`
2. Run the program and observe the behavior
3. Use `go tool pprof` to identify goroutine leaks
4. Implement your solution
5. Verify the fix using the provided tests

## Hints

1. Check the `fetchURL` function - is it properly handling all cases?
2. Look at how the context is being used
3. Consider what happens when errors occur
4. Think about channel closure timing

## Success Criteria

- No goroutine leaks (verify using pprof)
- All resources are properly cleaned up
- Error handling is robust
- The program maintains its concurrent behavior
- Tests pass successfully

## Bonus Challenge

1. Add timeout handling for slow URLs
2. Implement rate limiting
3. Add retry logic for failed requests
4. Create a more sophisticated error reporting system
