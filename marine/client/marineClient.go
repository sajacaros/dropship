package main

import (
	"context"
	"fmt"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
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

	client := marine.NewProjectServiceClient(cc)

	req := &marine.ProjectIdentity{
		Project: "command",
	}
	_, err = client.Stop(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC : %v", err)
	}
	log.Printf("complete")
}
