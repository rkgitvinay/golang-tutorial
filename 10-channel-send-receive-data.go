// Sending and Receiving Data

// Sending Data
// A goroutine sends data into a channel using the channel <- value syntax.

// Receiving Data
// The receiver retrieves data using <- channel.

package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <=5; i++ {
		ch <- i // Send data
		fmt.Println("Produced: ", i)
	}
	close(ch) // close the channle to indicate no more data
}

func consumer(ch chan int) {
	for val := range ch { //Receive data until channel is closed
		fmt.Println("Consumed:", val)
	}
}

func main() {
	ch := make(chan int)
	
	go producer(ch)
	consumer(ch)
}