package gateway

import (
	"context"
	"encoding/json"
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


	runtime.HTTPError = CustomHTTPError
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
		handlers.AllowedOrigins([]string{"*"}),
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

type errorBody struct {
	Code int32 `json:"code"`
	Err string `json:"message"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"message": "failed to marshal error message", "code":5000}`

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(http.StatusConflict)
	jErr := json.NewEncoder(w).Encode(errorBody{
		Err: err.Error(),
		Code: 5001,
	})

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}