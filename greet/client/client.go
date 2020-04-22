package main

import (
	"context"
	"log"

	"github.com/muzfr7/grpc-sample/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("Starting RPC client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	log.Printf("Starting to do Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Muzafar Ali",
			LastName:  "Jatoi",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
