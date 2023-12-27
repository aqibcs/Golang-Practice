package main

import (
	"fmt"
	"time"
)

func printNumbers(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	numbersChannel := make(chan int)
	go printNumbers(numbersChannel)

	for num := range numbersChannel {
		fmt.Println(num)
	}
}
