package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Option func(*http.Server) error

func WithAddr(addr string) Option {
	return func(s *http.Server) error {
		s.Addr = addr
		return nil
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(s *http.Server) error {
		if t > time.Second {
			return errors.New("read timeout value not allowed")
		}
		s.ReadTimeout = t
		return nil
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(s *http.Server) error {
		s.WriteTimeout = t
		return nil
	}
}

func NewServer(opts ...Option) (http.Server, error) {
	s := http.Server{}
	for _, opt := range opts {
		if err := opt(&s); err != nil {
			return s, err
		}
	}
	return s, nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	s, err := NewServer(
		WithAddr(":9191"),
		WithReadTimeout(30*time.Second),
		WithWriteTimeout(30*time.Millisecond),
	)

	if err != nil {
		log.Fatalln("Couldn't initialize server:", err)
	}

	s.Handler = mux

	log.Println("Server started", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
