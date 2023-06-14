package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		log.Fatalf("Error on create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error on response %v", err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
