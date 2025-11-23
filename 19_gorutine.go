package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
}
func main() {
	for i := 1; i <= 100; i++ {
		go task(i)
	}
	// Wait for user input to allow goroutines to finish
	time.Sleep(time.Second * 1) // wait for 1 second
}
