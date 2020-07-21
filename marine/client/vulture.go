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
	fmt.Println("Dropship Client start~")

	// length check
	args := os.Args[1:]
	argsLength := len(args)
	if argsLength != 1 || argsLength != 2  {
		log.Fatalf("need to 1 or 2 arguments, please try for 'vulture {command} ({project})'")
	}

	// init
	var req *marine.ProjectIdentity = nil
	if !strings.EqualFold(args[0], "install" ) {
		if argsLength != 2  {
			log.Fatalf("need to 2 arguments, please try for 'vulture {command} ({project})'")
		}
		req = &marine.ProjectIdentity{
			Project: args[1],
		}
	}

	// grpc init
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	client := marine.NewProjectServiceClient(cc)


	// execute
	if err := executeCommand(client, args[0], req); err !=nil {
		log.Fatalf("Error while calling RPC : %v", err)
	}

	log.Printf("complete")
}

func executeCommand(client marine.ProjectServiceClient, command string, req *marine.ProjectIdentity) error {
	var err error = nil
	if strings.EqualFold(command,"install") {
		_, err = client.Install(context.Background(), nil)
	} else if strings.EqualFold(command, "status") {
		_, err = client.Status(context.Background(), req)
	} else if strings.EqualFold(command, "start") {
		_, err = client.Start(context.Background(), req)
	} else if strings.EqualFold(command,"stop") {
		_, err = client.Stop(context.Background(), req)
	} else if strings.EqualFold(command,"update") {
		_, err = client.Update(context.Background(), req)
	} else {
		err = errors.New("not supported command")
	}
	return err
}
