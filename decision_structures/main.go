package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is true")
	} else {
		log.Println("isTrue is false")
	}

	myNum := 100
	isTrue = true

	if myNum > 99 && isTrue == true {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	}

	myVar := "Cat"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":

		log.Println("cat is set to dog")

	default:
		log.Println("Cat is something else")
	}
}
