//Synchronizing Goroutines with sync.WaitGroup

//The sync.WaitGroup is used to wait for a collection of goroutines to finish.
// You add the number of goroutines to wait for using Add(),
// decrement it with Done(), and block execution until all goroutines complete using Wait().

package main

import (
	"fmt"
	"sync"
	"time"
)

func workder(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d completed \n", id)
}


func main() {
	var wg sync.WaitGroup

	// Launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workder(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed!")
}