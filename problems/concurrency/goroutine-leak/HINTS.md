# Hints for Goroutine Leak Challenge

## Hint 1: Context Usage
The context is passed to `Start()` but it's not being used in the goroutines. Consider how you can use the context to cancel operations when the context is done.

## Hint 2: Channel Closure
The `results` channel is closed in a separate goroutine that waits for all fetches to complete. What happens if the context is cancelled before all fetches are done?

## Hint 3: Select Statement
Consider using a `select` statement in the fetch goroutines to handle both the context cancellation and the fetch operation. This allows you to respond to cancellation immediately.

## Hint 4: Error Handling
The current implementation doesn't handle the case where sending to the results channel might block. What happens if the receiver stops reading from the channel?

## Hint 5: Resource Cleanup
Think about how to ensure that all goroutines are properly cleaned up, even in error cases. Consider using a `defer` statement in the goroutines.

## Hint 6: Channel Buffer
The results channel is unbuffered. Would a buffered channel help prevent goroutine leaks in some cases?

## Hint 7: WaitGroup Usage
The WaitGroup is used to track goroutines, but it's not coordinated with context cancellation. How can you ensure the WaitGroup is properly decremented in all cases?

## Hint 8: Timeout Handling
The `fetchURL` function doesn't respect timeouts. Consider adding a timeout to the HTTP request using the context.

## Hint 9: Graceful Shutdown
Consider implementing a graceful shutdown mechanism that allows in-flight requests to complete while preventing new ones from starting.

## Hint 10: Final Solution Structure
A good solution should:
1. Use context for cancellation
2. Handle channel operations safely
3. Clean up resources properly
4. Handle timeouts gracefully
5. Prevent goroutine leaks in all scenarios
