//Simulate real-time data collection from sensors using goroutines and channels.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sensor(id int, data chan<- int) {
	for {
		reading := rand.Intn(100) // Simulate sensor reading
		data <- reading

		fmt.Printf("Sensor %d: Sent reading %d\n", id, reading)
		time.Sleep(time.Second)
	}
}

func monitor(data <-chan int) {
	for reading := range data {
		fmt.Printf("Monitor: Received reading %d\n", reading)
	}
}

func main() {
	data := make(chan int)

	// Launch two sensors
	go sensor(1, data)
	go sensor(2, data)

	// Monitor the data
	monitor(data)
}
