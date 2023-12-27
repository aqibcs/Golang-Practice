package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println("Element at index 0:", myArray[0])
	fmt.Println("Element at index 1:", myArray[1])
	fmt.Println("Element at index 2;", myArray[2])

	anotherArray := [3]int{4, 5, 6}

	fmt.Println("Element of anotherArray:", anotherArray)

	shortArray := [...]int{7, 8, 9}

	fmt.Println("Elements of shortArray:", shortArray)
}
