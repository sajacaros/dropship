package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("dropship-server-endpoint",  "localhost:50051", "dropship server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				OrigName:     true,
				EmitDefaults: true,
			},
		),
		)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := marine.RegisterProjectServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}


func main() {
	fmt.Printf("dropship gw start")
	//flag.Parse()
	//defer glog.Flush()

	if err := run(); err != nil {
		//glog.Fatal(err)
		log.Fatal("err for marine http")
	}
}