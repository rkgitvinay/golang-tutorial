// Goroutines are lightweight threads managed by the Go runtime.
// They enable concurrent execution of functions, making Go highly efficient for handling concurrent tasks.
// Unlike traditional threads, goroutines are inexpensive to create and execute,
// with smaller memory overhead and efficient scheduling by the Go runtime.

package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")

	fmt.Println("Main function running...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}
