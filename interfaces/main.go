package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	// Create a dog variable of type Dog.
	dog := Dog {
		Name: "Samson",
		Breed: "Gernan Shepered",
	}

	fmt.Println(fmt.Sprintf("the dog's name is %s and he is a %s.", dog.Name, dog.Breed))

	// Create gorilla variable of type Gorilla.
	gorilla := Gorilla {
		Name: "Gerald",
		Color: "black",
		NumberOfTeeth: 32,
	}

	fmt.Println(fmt.Sprintf("The gorilla's name is %s. She has %d teeth and she is %s in color.",
	gorilla.Name, gorilla.NumberOfTeeth, gorilla.Color))
}
