package process

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() error {
	log.Println("grpc server serve, gateway://localhost:50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	marine.RegisterProjectServiceServer(s, &dropship{})

	err = s.Serve(lis)
	log.Printf("grpc server terminate. %v\n", err)
	return err
}

type dropship struct {}

func (*dropship) Summary(ctx context.Context, _ *empty.Empty) (*marine.StatusSummary, error) {
	return Summary()
}
func (*dropship) Status(ctx context.Context, request *marine.ProjectIdentity) (*marine.ProjectStatus, error){
	return Status(request.Project), nil
}
func (*dropship) Start(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("start %s", request.Project)
	err := Start(request.Project)
	return &empty.Empty{}, err
}
func (*dropship) Stop(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("stop %s", request.Project)
	err := Stop(request.Project)
	return &empty.Empty{}, err
}
func (*dropship) Update(ctx context.Context, request *marine.ProjectIdentity) (*empty.Empty, error){
	log.Printf("update %s", request.Project)
	err := Update(request.Project)
	return &empty.Empty{}, err
}

func (*dropship) Install(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	log.Printf("Install")
	err := Install()
	return &empty.Empty{}, err
}

func (*dropship) Dependency( ctx context.Context, in *empty.Empty) (*marine.ProjectDependency, error) {
	return Dependency()
}