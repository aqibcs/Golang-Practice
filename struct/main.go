package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Aqib",
		LastName:    "Shah",
		PhoneNumber: "0090909009",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
