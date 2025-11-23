package main

import (
	"fmt"
	"time"
)

func processvalue(ch chan string) {

	fmt.Println(<-ch)

}
func main() {
	// Channels are used to communicate between goroutines
	// They provide a way for one goroutine to send data to another goroutine

	// Create a channel of type int
	// message := make(chan string)

	// // Start a goroutine that sends a message to the channel
	// message <- "Hello from goroutine!" // blocking operation

	// msg := <-message
	// fmt.Println(msg)

	var message chan string
	message = make(chan string)
	go processvalue(message)
	message <- "Hello from main"

	time.Sleep(time.Second * 1)

}
