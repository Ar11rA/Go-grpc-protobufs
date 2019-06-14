package main

import (
	greetpbServerStream "go/protobufs/ServerStreaming/greetpbServerStreaming"
	"time"

	"log"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct{}

// GreetStream ...
func (*Server) GreetStream(req *greetpbServerStream.GreetRequest, stream greetpbServerStream.GreetService_GreetStreamServer) error {
	log.Printf("Greet function invoked with req %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	greeting := "Hello " + firstName + " " + lastName
	for _, letter := range greeting {
		greetResponse := &greetpbServerStream.GreetResponse{
			Result: string(letter),
		}
		stream.Send(greetResponse)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
