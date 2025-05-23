# Interface Satisfaction & Type Assertion Challenge

## Problem Description

You are given a plugin system where different types implement various interfaces. The current implementation has subtle bugs related to interface satisfaction and type assertions. Your task is to:

1. Identify and fix interface satisfaction issues
2. Correctly use type assertions and type switches
3. Handle nil interface values and assertion panics
4. Refactor the code for idiomatic Go interface usage

## Learning Objectives

- Understanding interface satisfaction in Go
- Using type assertions and type switches safely
- Handling nil interfaces and assertion panics
- Idiomatic interface design and usage
- Debugging interface-related bugs

## Getting Started

1. Read through the code in `main.go`
2. Run the program and observe the output and panics
3. Identify where type assertions or interface satisfaction fail
4. Implement your solution
5. Verify the fix using the provided tests

## Hints

1. Check if all types actually implement the interfaces they claim
2. Use type switches to handle multiple interface types
3. Be careful with nil interface values (typed nil vs untyped nil)
4. Use the `ok` idiom for safe type assertions
5. Consider interface embedding for extensibility

## Success Criteria

- No panics from failed type assertions
- All interface methods are satisfied
- Idiomatic and robust interface usage
- Tests pass successfully

## Bonus Challenge

1. Implement a plugin registry using interfaces
2. Add dynamic plugin loading with reflection
3. Create mock implementations for testing
4. Add interface-based middleware for plugins
