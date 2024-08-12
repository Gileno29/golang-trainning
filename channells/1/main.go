package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador("a")

	go contador("b") //inicia uma thread
	time.Sleep(time.Second * 10)
	//fmt.Println("Hello 1")
	//fmt.Println("Hello 2")

	//time.Sleep(time.Second)
	//fmt.Printf("fim...")
}
