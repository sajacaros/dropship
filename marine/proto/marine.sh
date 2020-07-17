# grpc and gateway
protoc -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. marine/proto/marine.proto

protoc -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:build/gen --grpc-gateway_out=logtostderr=true:build/gen marine/proto/marine.proto

set GOOS=linux&&set GOARCH=amd64&&go build -o build/out/marineServer marine/server/marineServer.go
