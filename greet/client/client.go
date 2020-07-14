package main

import (
	"context"
	"fmt"
	"github.com/sajacaros/dropship/greet/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client start~")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := proto.NewGreetServiceClient(cc)

	req := &proto.GreetRequest{
		Greeting: &proto.Greeting{
			FirstName: "hello",
			LastName: "world",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC : %v", err)
	}
	log.Printf("Response from Gree: %v", res.Result)
}
