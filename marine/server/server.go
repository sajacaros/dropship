package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

type dropship struct {}

func (*dropship)Summary(ctx context.Context, request *empty.Empty) (*marine.StatusSummary, error) {
	log.Printf("summary")
	return nil, nil
}
func (*dropship)Status(ctx context.Context, request *marine.ProjectIdentity) (*marine.ProjectStatus, error){
	log.Printf("status %s", request.Project)
	return nil, nil
}
func (*dropship)Start(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("start %s", request.Project)
	return nil, nil
}
func (*dropship)Stop(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("stop %s", request.Project)
	return nil, nil
}
func (*dropship)Update(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("update %s", request.Project)
	return nil, nil
}

//func (*server)Greet(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
//	if request == nil || request.Greeting == nil {
//		return nil, errors.New("request failed")
//	}
//	firstName := request.Greeting.FirstName
//	result := "Hello " + firstName
//	res := &proto.GreetResponse{
//		Result: result,
//	}
//	return res, nil
//}

func main() {
	fmt.Println("Start Dropship Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	marine.RegisterProjectServiceServer(s, &dropship{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
