package main

import (
	greetpbClientStream "go/protobufs/ClientStreaming/greetpbClientStreaming"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct{}

// GreetStream ...
func (*Server) GreetStream(stream greetpbClientStream.GreetService_GreetStreamServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpbClientStream.GreetResponse{
				Result: result,
			})
		}
		firstName := req.GetGreeting().GetFirstName()
		log.Println("Received from server: ", firstName)
		result += "ok " + firstName + " "
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("No server available! %v", err)
	}
	server := grpc.NewServer()
	greetpbClientStream.RegisterGreetServiceServer(server, &Server{})

	log.Println("Starting server at 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
