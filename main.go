package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Option func(*http.Server)

func WithAddr(addr string) Option {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(s *http.Server) {
		s.WriteTimeout = d
	}
}

func NewServer(opts ...Option) http.Server {
	s := http.Server{}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	s := NewServer(
		WithAddr(":9191"),
		WithReadTimeout(1*time.Millisecond),
		WithWriteTimeout(1*time.Millisecond),
	)
	s.Handler = mux

	log.Println("Server started", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
