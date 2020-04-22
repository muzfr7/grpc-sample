package main

import (
	"context"
	"log"

	"github.com/muzfr7/grpc-sample/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("Starting RPC client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do Unary RPC...")

	req := &calculatorpb.CalculateRequest{
		Calculator: &calculatorpb.Calculator{
			Num1: 3,
			Num2: 10,
		},
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculate RPC: %v", err)
	}

	log.Printf("Response from Calculate: %v", res.Result)
}
