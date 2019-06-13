package main

import (
	"context"
	"go/protobufs/Unary/greetpbUnary"
	"log"

	"google.golang.org/grpc"
)

func getGreeting(client greetpbUnary.GreetServiceClient) {
	request := &greetpbUnary.GreetRequest{
		Greeting: &greetpbUnary.Greeting{
			FirstName: "firstName",
			LastName:  "lastName",
		},
	}
	response, err := client.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Failed to call greet %v", err)
	}

	log.Printf("Response from greet %v", response.Result)
}

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client doesn't connect!! %v", err)
	}

	defer conn.Close()

	client := greetpbUnary.NewGreetServiceClient(conn)
	log.Println("Client running", client)
	getGreeting(client)
}
