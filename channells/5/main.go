package main

import (
	"fmt"
	"time"
)

// Se a variavel queue não for utilizada a go routinet vai ficar presa e não sera executada
func main() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}
