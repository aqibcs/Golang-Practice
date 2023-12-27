package main

import "fmt"

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("End") // This won't be executed
}

func main() {
	recoverExample()
	fmt.Println("Continuing after recoverExample")
}
