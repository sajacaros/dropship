# grpc and gateway
protoc -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. greet/proto/greet.proto
