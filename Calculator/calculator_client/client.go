package main

import (
	"context"
	"go/protobufs/Calculator/calculatorpb"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func doCalculations(client calculatorpb.CalculatorServiceClient, number1 string, number2 string) {
	num1, _ := strconv.ParseInt(number1, 10, 32)
	num2, _ := strconv.ParseInt(number2, 10, 32)
	request := &calculatorpb.CalculationRequest{
		FirstNumber:  int32(num1),
		SecondNumber: int32(num2),
	}
	response, err := client.Ops(context.Background(), request)
	if err != nil {
		log.Fatalf("Not able to perform Ops!")
	}
	log.Printf("%v", response)
}

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client doesn't connect!! %v", err)
	}

	defer conn.Close()
	client := calculatorpb.NewCalculatorServiceClient(conn)
	log.Println("Client running", client)
	doCalculations(client, os.Args[1], os.Args[2])
}
