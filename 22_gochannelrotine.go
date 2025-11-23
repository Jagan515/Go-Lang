package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(msg chan int) {
	for m := range msg {
		fmt.Println("Received:", m)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	message := make(chan int)
	go worker(message)

	for {
		message <- rand.Intn(100)
	}
}
