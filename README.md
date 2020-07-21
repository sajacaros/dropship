"# dropship"

## marine
```
------                         --------
|http|     --(grpc unary)->    |server|
------                         --------
--------                       --------
|client|   --(grpc unary)->    |server|
--------                       -------- 
```

## grpc, grpc-gw
    ```
    protoc -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:build/gen --grpc-gateway_out=logtostderr=true:build/gen marine/proto/marine.proto
    ```
## win -> linux binary
    ```
    set GOOS=linux&&set GOARCH=amd64&&go build -o build/out/marineServer marine/server/marineServer.go
    set GOOS=linux&&set GOARCH=amd64&&go build -o build/out/vulture marine/client/vulture.go
    set GOOS=linux&&set GOARCH=amd64&&go build -o build/out/marineHttp marine/http/marineHttp.go
    ```