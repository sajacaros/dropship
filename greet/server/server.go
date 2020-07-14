package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sajacaros/dropship/greet/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server)Greet(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	if request == nil || request.Greeting == nil {
		return nil, errors.New("request failed")
	}
	firstName := request.Greeting.FirstName
	result := "Hello " + firstName
	res := &proto.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Start Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
