package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//this is the general mode of declaration maps
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"])

	me := User{
		FirstName: "Gileno",
		LastName:  "Duarte",
	}

	myMap2 := make(map[string]User)

	myMap2["me"] = me
	log.Println(myMap2["me"].FirstName)
}
