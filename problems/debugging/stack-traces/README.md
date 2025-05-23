# Stack Trace and Debugging Challenge

## Problem Description

You are given a complex service that processes user data and occasionally crashes with cryptic error messages. Your task is to:

1. Identify the root cause of the crashes
2. Implement proper error handling and logging
3. Add meaningful stack traces
4. Create a debugging strategy
5. Implement recovery mechanisms

## Learning Objectives

- Understanding Go's error handling patterns
- Working with stack traces and panic recovery
- Implementing proper logging
- Debugging techniques and tools
- Error wrapping and context preservation

## Getting Started

1. Read through the code in `main.go`
2. Run the program and observe the crashes
3. Use the debugger to trace the execution
4. Implement your solution
5. Verify the fix using the provided tests

## Hints

1. Look at how errors are propagated
2. Consider using `runtime.Caller` for stack traces
3. Think about error wrapping with context
4. Consider implementing a custom error type

## Success Criteria

- No unexpected crashes
- Meaningful error messages with context
- Proper stack traces in logs
- Graceful error recovery
- Tests pass successfully

## Bonus Challenge

1. Implement structured logging
2. Add metrics for error tracking
3. Create a custom debugger
4. Implement error reporting to external services
