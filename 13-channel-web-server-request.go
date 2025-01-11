//Use a buffered channel to handle incoming requests efficiently.

package main

import (
	"fmt"
	"time"
)

func handleRequest(id int, ch chan int) {
	for req := range ch {
		fmt.Printf("Handler %d processing request %d\n", id, req)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
	}
}

func main() {
	ch := make(chan int, 5) // Buffered channel

	// launch req handlers
	for i := 1; i <= 3; i++ {
		go handleRequest(i, ch)
	}

	// Simulate incoming requests
	for i := 1; i <= 10; i++ {
		fmt.Printf("Incoming request: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	time.Sleep(2 * time.Second) // Allow handlers to complete
}
