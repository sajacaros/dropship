package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"github.com/sajacaros/dropship/marine/process"
	"google.golang.org/grpc"
	"log"
	"net"
)

type dropship struct {}

func (*dropship) Summary(ctx context.Context, request *empty.Empty) (*marine.StatusSummary, error) {
	log.Printf("summary")
	return nil, nil
}
func (*dropship) Status(ctx context.Context, request *marine.ProjectIdentity) (*marine.ProjectStatus, error){
	log.Printf("status %s", request.Project)
	return nil, nil
}
func (*dropship) Start(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("start %s", request.Project)
	err := process.Start(request.Project)
	return &empty.Empty{}, err
}
func (*dropship) Stop(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("stop %s", request.Project)
	err := process.Stop(request.Project)
	return &empty.Empty{}, err
}
func (*dropship) Update(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("update %s", request.Project)
	err := process.Update(request.Project)
	return &empty.Empty{}, err
}

func (*dropship) Install(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	log.Printf("Install")
	err := process.Install()
	return &empty.Empty{}, err
}

func main() {
	log.Println("Start Dropship Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	marine.RegisterProjectServiceServer(s, &dropship{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
