package main

import (
	"github.com/sajacaros/dropship/marine/gateway"
	"github.com/sajacaros/dropship/marine/process"
	"github.com/sajacaros/dropship/marine/static"
	"log"
	"sync"
)

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
	gateway.Run()
	waitGroup.Done()
}

func serveGRpcServer(waitGroup *sync.WaitGroup)  {
	process.Run()
	waitGroup.Done()
}

func serveStaticServer(waitGroup *sync.WaitGroup) {
	static.Run()
	waitGroup.Done()
}