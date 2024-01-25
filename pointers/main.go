package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to: ", myString)
	//use & for reference memory address
	changeUsingPointer(&myString)
	log.Println("myString is set to: ", myString)

}

//use * to pointer its self to make a pointer
func changeUsingPointer(s *string) {
	log.Println("teh value of s is: ", s)
	newValue := "Red"
	*s = newValue
}
