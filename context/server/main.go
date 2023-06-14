package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Start Request")
	defer log.Println("End Request")

	select {
	case <-time.After(time.Second * 5):
		log.Println("request sucessfull")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request sucessfull"))

	case <-ctx.Done():

		log.Println("Request Cancel")
		http.Error(w, "Request Cancelada", http.StatusRequestTimeout)
	}

}
