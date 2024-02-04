package main

import "log"

func main() {
	//this is the general mode of declaration maps
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"])
}
