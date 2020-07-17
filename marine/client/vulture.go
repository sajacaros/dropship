package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("2 arguments, marineClient {project} {command}")
	}

	fmt.Println("Dropship Client start~")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := marine.NewProjectServiceClient(cc)

	req := &marine.ProjectIdentity{
		Project: args[0],
	}

	if err := executeCommand(client, args[1], req); err !=nil {
		log.Fatalf("Error while calling RPC : %v", err)
	}

	log.Printf("complete")
}

func executeCommand(client marine.ProjectServiceClient, input string, req *marine.ProjectIdentity) error {
	command := strings.ToLower(input)
	var err error = nil
	if command == "status" {
		_, err = client.Status(context.Background(), req)
	} else if command == "start" {
		_, err = client.Start(context.Background(), req)
	} else if command == "stop" {
		_, err = client.Stop(context.Background(), req)
	} else if command == "update" {
		_, err = client.Update(context.Background(), req)
	} else {
		err = errors.New("not supported command")
	}
	return err
}
