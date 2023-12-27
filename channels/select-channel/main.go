package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: sends a message to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Goroutine 1"
	}()

	// Goroutine 2: sends a message to ch2 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from Goroutine 2"
	}()

	// Use select to wait for either ch1 or ch2
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
