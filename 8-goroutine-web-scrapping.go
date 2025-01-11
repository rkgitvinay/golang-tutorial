//Web Scraping

package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Fetching data from %s\n", url)
	time.Sleep(time.Second * 3)
	fmt.Printf("Data fetched from %s\n", url)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{"https://site1.com", "https://site2.com", "https://site3.com"}

	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}

	wg.Wait()

	fmt.Println("All data fetched")
}
