package gateway

import (
	"context"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"google.golang.org/grpc"
	"net/http"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("dropship-server-endpoint",  "localhost:50051", "dropship server endpoint")
)

func Run() error {
	flag.Parse()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	basicMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				OrigName:     true,
				EmitDefaults: true,
			},
		),
		)
	mux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"content-type", "x-foobar-key"}),
	)(basicMux)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := marine.RegisterProjectServiceHandlerFromEndpoint(ctx, basicMux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":3001", mux)
}


//func main() {
//	fmt.Printf("dropship gw start")
//
//	if err := run(); err != nil {
//		log.Fatal("err for marine gateway")
//	}
//}