package main

import (
	"fmt"
	"time"
)

func countDown(n int, done chan bool) {
	for n >= 0 {
		fmt.Println(n)
		n--
		time.Sleep(time.Millisecond * 500)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go countDown(5, done)

	// Block until the countdown goroutine signals it's done
	<-done
}
