package main

import (
	"go/protobufs/example2/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct{}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("No server available! %v", err)
	}
	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
