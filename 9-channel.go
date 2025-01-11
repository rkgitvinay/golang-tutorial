// Channels are a way to communicate between goroutines.
// They allow one goroutine to send data to another goroutine,
// ensuring safe and synchronized data exchange.

package main

import "fmt"

func main() {

	// Unbuffered Channels
    // Data is sent and received synchronously.
    // The sender is blocked until the receiver receives the data, and vice versa.

	u_ch := make(chan string)

	go func ()  {
		u_ch <- "Hello Unbuffered channel"
	}()
	msg := <-u_ch
	fmt.Println(msg)


	// Buffered Channels
    // Data is sent asynchronously up to the capacity of the buffer.
    // The sender is only blocked when the buffer is full, and the receiver is blocked when the buffer is empty.

	b_ch := make(chan string, 2)
	b_ch <- "Message 1"
	b_ch <- "Message 2"

	fmt.Println(<-b_ch)
	fmt.Println(<-b_ch)
}