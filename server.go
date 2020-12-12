package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}

}

func (s *Server) Listen() error {
	fmt.Println("Server listening on port", s.port[1:len(s.port)])
	return http.ListenAndServe(s.port, nil)
}
