package main

import (
	"log"
)

func main() {
	// The if statement
	var isTrue bool

	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat2"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 10
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and istrue is set to true")
	} else if myNum < 100 && !isTrue {
		log.Println("myNum is less than 100 and isTrue is set to false")
	}

	// The switch statement
	myVar := "fish"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "horse":
		log.Println("myVar is set to horse")
	default:
		log.Println("myVar is something else")
	}
}
