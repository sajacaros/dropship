package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Dropship Client start...")

	// length check
	args := os.Args[1:]
	argsLength := len(args)
	if !(argsLength == 1 || argsLength == 2) {
		log.Fatalf("need to 1 or 2 arguments, please try for 'vulture {command} ({project})'")
	}

	// init
	var req *marine.ProjectIdentity = nil
	if !(strings.EqualFold(args[0], "install") || strings.EqualFold(args[0], "summary")) {
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
		_, err = client.Install(context.Background(), &empty.Empty{})
	} else if strings.EqualFold(command, "status") {
		var status *marine.ProjectStatus
		status, err = client.Status(context.Background(), req)
		log.Println(statusToString(status))
	} else if strings.EqualFold(command, "start") {
		_, err = client.Start(context.Background(), req)
	} else if strings.EqualFold(command, "stop") {
		_, err = client.Stop(context.Background(), req)
	} else if strings.EqualFold(command, "update") {
		_, err = client.Update(context.Background(), req)
	} else if strings.EqualFold(command, "summary"){
		var summary *marine.StatusSummary
		summary, err = client.Summary(context.Background(), &empty.Empty{})
		for _, status := range summary.Projects {
			log.Println(statusToString(status))
		}
	} else {
		err = errors.New("not supported command")
	}
	return err
}

func statusToString(status *marine.ProjectStatus) string {
	var sb strings.Builder
	sb.WriteString(status.Project)
	sb.WriteString(" status : ")
	sb.WriteString(status.Status)
	if status.Pid != 0 {
		sb.WriteString("(pid : ")
		sb.WriteString(strconv.Itoa(int(status.Pid)))
		sb.WriteString(", uptime : ")
		sb.WriteString(status.Uptime)
		sb.WriteString(")")
	}
	return sb.String()
}

