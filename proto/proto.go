package proto

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, message *Request) (*Response, error) {
	log.Printf("Received message from client: %s", message.Name)
	return &Response{Greeting: "Hello from the server"}, nil
}
