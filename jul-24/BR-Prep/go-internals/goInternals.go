/*
Question 1

Describe how Go handles error handling with the use of multiple return values and the error
interface. What are the benefits of this approach?

Question 2

How can you prevent race conditions in Go? Explain different techniques or synchronization
primitives available in Go to ensure safe concurrent access to shared resources.

Race conditions occur when multiple goroutines access shared resources concurrently and at least one of the accesses is a write. This can lead to unpredictable behavior and bugs that are often hard to reproduce and fix. Go provides several synchronization primitives to prevent race conditions and ensure safe concurrent access to shared resources.

Techniques to Prevent Race Conditions
Mutexes
RWMutexes
Channels
Atomic Operations
WaitGroup
Once

Question 3

In a concurrent scenario, how can you detect and handle deadlocks in Go? What strategies
or techniques can be employed to avoid or mitigate deadlocks?
Question 4

How can middleware be used in a Go web framework, and what are some common use cases?

Ans:
Logging
Authentication & Autherization
Error Handling
CROS - Cross Origin Resource Sharing
RAte Limiting

Question 5

How does Go's interface system work internally, and what are some advantages and limitations
of using interfaces in Go?

Question 6

How does Go manage dependencies with its module system, and what are some best practices
for versioning and dependency management in Go projects?

go mod init example.com/myproject
go get example.com/somepackage@v1.2.3



Question 7

Describe the memory layout of Go programs and how it affects performance, including the stack,
heap, and garbage collection considerations.
*/

