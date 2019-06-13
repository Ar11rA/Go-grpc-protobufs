package main

import (
	"context"
	"go/protobufs/Calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct{}

// Ops ...
func (*Server) Ops(ctx context.Context, req *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error) {
	number1 := req.GetFirstNumber()
	number2 := req.GetSecondNumber()
	calculationResponse := &calculatorpb.CalculationResponse{
		AddResult:      number1 + number2,
		SubtractResult: number1 - number2,
		DivideResult:   (int32)(number1 / number2),
		MultiplyResult: number1 * number2,
	}
	return calculationResponse, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("No server available! %v", err)
	}
	server := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(server, &Server{})
	log.Println("Starting server at 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
