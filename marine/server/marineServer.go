package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"github.com/sajacaros/dropship/marine/gateway"
	"github.com/sajacaros/dropship/marine/process"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

type dropship struct {}

func (*dropship) Summary(ctx context.Context, _ *empty.Empty) (*marine.StatusSummary, error) {
	log.Printf("summary")
	return process.Summary()
}
func (*dropship) Status(ctx context.Context, request *marine.ProjectIdentity) (*marine.ProjectStatus, error){
	log.Printf("status %s", request.Project)
	return process.Status(request.Project), nil
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

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	go serveGRpcServer(&waitGroup)
	go serveStaticServer(&waitGroup)
	go serveGRpcGatewayServer(&waitGroup)

	waitGroup.Wait()
}

func serveGRpcGatewayServer(waitGroup *sync.WaitGroup) {
	log.Println("grpc-gateway server serve")

	if err := gateway.Run(); err != nil {
		waitGroup.Done()
		log.Printf("grpc-gateway server terminate. %v\n", err)
	}
}

func serveGRpcServer(waitGroup *sync.WaitGroup)  {
	log.Println("grpc server serve, gateway://localhost:50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	marine.RegisterProjectServiceServer(s, &dropship{})

	if err := s.Serve(lis); err != nil {
		waitGroup.Done()
		log.Printf("grpc server terminate. %v\n", err)
	}
}

func serveStaticServer(waitGroup *sync.WaitGroup) {
	home, _ := os.UserHomeDir()
	workingDir := home + "/workspace/dropship"
	log.Println("static server serve, gateway://localhost:3000")
	fs := http.FileServer(http.Dir(workingDir+"/static/index.html"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		waitGroup.Done()
		log.Printf("static server terminate. %v\n", err)
	}
}