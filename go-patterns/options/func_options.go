package main

import (
	"fmt"
	"time"
)

func main() {
	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer("localhost", 8080, Protocol("tcp"), Timeout(300*time.Second))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

// a type of option function
// a few functions that accept Server as argument and return an option function
// a NewServer to create server with options

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	//Maxconns int
	//TLS      *tls.Config
}

func NewServer(address string, port int, options ...Option) (*Server, error) {
	s := &Server{Addr: address, Port: port}
	for _, option := range options {
		option(s)
	}
	return s, nil
}
