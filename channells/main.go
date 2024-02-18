package main

import "github.com/Gileno29/golang-trainning/channells/helpers"

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)
	num := <-intChan

	println(num)
}
