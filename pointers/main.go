package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to: ", myString)
	changeUsingPointer(&myString)
	log.Println("myString is set to: ", myString)

}

func changeUsingPointer(s *string) {
	log.Println("teh value of s is: ", s)
	newValue := "Red"
	*s = newValue
}