package main

import (
	"fmt"
	//"time"
)

// func emailsender(emailchan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailchan {
// 		fmt.Println("sending email to", email)
// 	}
// }

// func main() {
// 	emailchan := make(chan string, 100)
// 	done := make(chan bool)

// 	go emailsender(emailchan, done)

// 	for i := 0; i < 100; i++ {
// 		emailchan <- fmt.Sprintf("%d@gmail.com", i)
// 		time.Sleep(time.Second)
// 	}

// 	close(emailchan) // IMPORTANT

// 	<-done
// 	fmt.Println("all emails sent")
// }

// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func() {
// 		chan1 <- 10
// 	}()

// 	go func() {
// 		chan2 <- "pong"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-chan1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-chan2:
// 			fmt.Println("received", msg2)
// 		}
// 	}
// }

//  send only channel & receive only channel

func emailsender(emailchan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailchan {
		fmt.Println("sending email to", email)
	}
}

func main() {
	emailchan := make(chan string, 100)
	done := make(chan bool)

	go emailsender(emailchan, done)

	for i := 0; i < 100; i++ {
		emailchan <- fmt.Sprintf("%d@gmail.com", i)
		time.Sleep(time.Second)
	}

	close(emailchan) // IMPORTANT

	<-done
	fmt.Println("all emails sent")
}
