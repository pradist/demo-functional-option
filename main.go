package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewServer(addr string, readT, writeT time.Duration) http.Server {
	s := http.Server{
		Addr:         addr,
		ReadTimeout:  readT,
		WriteTimeout: writeT,
	}

	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	s := NewServer(":8080", 1*time.Second, 1*time.Second)
	s.Handler = mux

	log.Println("Server started", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
