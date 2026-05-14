# Goroutines - Concurrency in Go

Concurrency is the ability of a program to handle multiple tasks at the same time. This is Go’s "killer feature," as it was designed from the ground up to make high-performance, multi-threaded applications easy to write.

1. The Problem: Sequential Execution
By default, Go programs run sequentially. If one function (like sending an email or generating a PDF) takes 10 seconds to complete, the entire program "blocks" and waits. The next line of code cannot run until the previous one is finished.

• The Result: Users have to wait, and the application feels slow or unresponsive.


2. The Solution: Goroutines (go keyword)
A Goroutine is a lightweight thread managed by the Go runtime.

• How to use it: Simply add the keyword go before a function call.
Example: go sendTicket(userTickets, firstName, lastName, email)

• What happens? Instead of waiting for sendTicket to finish, Go spins off a new "green thread" to handle that function in the background and immediately moves to the next line in the main program.


# Managing Threads with WaitGroups

If the main thread (the main function) finishes before the background Goroutines are done, the program will exit and kill those background tasks. To prevent this, you use a WaitGroup from the sync package.

• wg.Add(int): Sets the number of Goroutines the main thread should wait for.

• wg.Wait(): Tells the main thread to stop and wait until the counter hits zero.

• wg.Done(): Called inside the Goroutine to signify that its work is finished, decreasing the counter by one.


# Why Go's Concurrency is Better

Go routines are significantly more efficient than traditional threads in languages like Java or C++.

| Feature | OS Threads (Java/C++) | Goroutines (Go) |
| Memory | High (MBs per thread) | Very Low (KBs per thread) |
| Startup | Slower (OS overhead) | Extremely Fast |
| Scale | Hundreds of threads | Millions of Goroutines |
| Model | Managed by OS | Managed by Go Runtime |


# Channels

Go provides Channels as a built-in way for Goroutines to talk to each other safely. This avoids "race conditions" where two threads try to update the same data at once.