package main

import (
	"context"
	greetpbServerStream "go/protobufs/ServerStreaming/greetpbServerStreaming"
	"io"
	"log"

	"google.golang.org/grpc"
)

func getGreetStream(client greetpbServerStream.GreetServiceClient) {
	request := &greetpbServerStream.GreetRequest{
		Greeting: &greetpbServerStream.Greeting{
			FirstName: "firstName",
			LastName:  "lastName",
		},
	}
	stream, err := client.GreetStream(context.Background(), request)

	if err != nil {
		log.Fatalf("Failed to call greet %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive stream")
		}
		log.Printf("Response from stream %v", msg.GetResult())
	}
}

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client doesn't connect!! %v", err)
	}

	defer conn.Close()

	client := greetpbServerStream.NewGreetServiceClient(conn)
	log.Println("Client running", client)
	getGreetStream(client)
}
