package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Task %d is running\n", id)

}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go task(i, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish till goroutine counter is zero
}
