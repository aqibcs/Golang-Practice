package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Maps
	myMap := make(map[string]User)

	me := User{
		FirstName: "Aqib",
		LastName:  "Shah",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	// Slices
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)

	// Shorthand for slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// Print first two elements of slices, starting at first position
	log.Println(numbers[0:2])

	// Create a slice of strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
