package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type ServerOpts struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServerOpts() ServerOpts {
	return ServerOpts{
		Address:      ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}

func NewServer(opt ServerOpts) http.Server {
	s := http.Server{
		Addr:         opt.Address,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
	}
	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	s := NewServer(NewServerOpts())
	s.Handler = mux

	log.Println("Server started", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
