// Task Distribution System
// Distribute tasks to multiple workers using buffered channels.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
	}
}

func main() {
	tasks := make(chan int, 10) // buffered channel

	var wg sync.WaitGroup

	// Start worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// send tasks
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks) // Close the channel to signal no more tasks

	wg.Wait()

	fmt.Println("All tasks completed")
}
