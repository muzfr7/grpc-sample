package main

import (
	"context"
	"log"
	"net"

	"github.com/muzfr7/grpc-sample/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v", req)

	firstName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName

	return &greetpb.GreetResponse{
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
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
