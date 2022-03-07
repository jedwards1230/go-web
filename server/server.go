package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	chat "github.com/jedwards1230/go-web/proto"

	"google.golang.org/grpc"
)

func main() {
	// Parse cli arguments
	ip := flag.String("ip", "127.0.0.1", "IP for the client to connect to")
	port := flag.Int("port", 8090, "Port for the server to listen on")
	flag.Parse()

	// Start http server
	log.Println(fmt.Sprintf("Starting server at %v:%d", *ip, *port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// Establish gRPC server
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC")
	}
}
