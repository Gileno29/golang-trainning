package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//good pratice
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()
	bookHotel(ctx)
}

//is a good pratice use ctx as the first param in the function when we are using that
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("time exceed")

	case <-time.After(time.Second * 5):
		fmt.Println("Reserved!!!")
	}

}
