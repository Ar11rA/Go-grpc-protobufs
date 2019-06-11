package main

import (
	"fmt"
	"go/protobufs/example2/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:51051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client doesn't connect!! %v", err)
	}
	defer conn.Close()
	client := greetpb.NewGreetServiceClient(conn)
	fmt.Println("Client running", client)
}
