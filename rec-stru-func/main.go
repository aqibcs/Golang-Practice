package main

import (
	"fmt"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct

	myVar.FirstName = "Aqib"

	myVar2 := myStruct {
		FirstName: "Asad",
	}
	fmt.Println("myVar is set to", myVar.printFirstName())
	fmt.Println("myVar is set to", myVar2.printFirstName())
}