package main

import (
	"context"
	greetpbServerStream "go/protobufs/ServerStreaming/greetpbServerStreaming"

	"log"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct{}

// Greet ...
func (*Server) Greet(ctx context.Context, req *greetpbServerStream.GreetRequest) (*greetpbServerStream.GreetResponse, error) {
	log.Printf("Greet function invoked with req %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	greeting := "Hello " + firstName + " " + lastName
	greetResponse := &greetpbServerStream.GreetResponse{
		Result: greeting,
	}
	return greetResponse, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("No server available! %v", err)
	}
	server := grpc.NewServer()
	greetpbServerStream.RegisterGreetServiceServer(server, &Server{})

	log.Println("Starting server at 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
