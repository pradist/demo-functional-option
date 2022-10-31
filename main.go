package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Server started", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
