// chanels fazem com que threads se comuniquem(compartilhem variaveis)
package main

import "fmt"

// T1
func main() {
	hello := make(chan string)

	// T2
	go func() {
		hello <- "Hello world"
	}()

	result := <-hello

	fmt.Println(result)
}
