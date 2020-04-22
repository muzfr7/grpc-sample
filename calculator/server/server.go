package main

import (
	"context"
	"log"
	"net"

	"github.com/muzfr7/grpc-sample/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	log.Printf("Calculate function was invoked with: %v", req)

	num1 := req.GetCalculator().GetNum1()
	num2 := req.GetCalculator().GetNum2()

	result := num1 + num2

	return &calculatorpb.CalculateResponse{
		Result: result,
	}, nil
}

func main() {
	log.Printf("Starting RPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
