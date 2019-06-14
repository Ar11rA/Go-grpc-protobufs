package main

import (
	"context"
	greetpbClientStream "go/protobufs/ClientStreaming/greetpbClientStreaming"
	"log"
	"time"

	"google.golang.org/grpc"
)

func getGreetStream(client greetpbClientStream.GreetServiceClient) {
	requests := []*greetpbClientStream.GreetRequest{
		&greetpbClientStream.GreetRequest{
			Greeting: &greetpbClientStream.Greeting{
				FirstName: "Hola",
			},
		},
		&greetpbClientStream.GreetRequest{
			Greeting: &greetpbClientStream.Greeting{
				FirstName: "Hello",
			},
		},
		&greetpbClientStream.GreetRequest{
			Greeting: &greetpbClientStream.Greeting{
				FirstName: "Hi",
			},
		},
	}
	stream, _ := client.GreetStream(context.Background())
	for _, request := range requests {
		stream.Send(request)
		time.Sleep(1 * time.Second)
	}
	res, _ := stream.CloseAndRecv()
	log.Printf("Final response %v", res)
}

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client doesn't connect!! %v", err)
	}

	defer conn.Close()

	client := greetpbClientStream.NewGreetServiceClient(conn)
	log.Println("Client running", client)
	getGreetStream(client)
}
