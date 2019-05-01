package main

import (
	"fmt"
	"log"
	"net"
	"os"

	sample "github.com/tommy-sho/memory-store-sample/proto"
	server "github.com/tommy-sho/memory-store-sample/server/internal"
	"google.golang.org/grpc"
)

func main() {
	sampleServer, err := server.NewsampleServer(os.Getenv("REDIS_HOST"))
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", "10000"))
	if err != nil {
		log.Fatal("Command: failed to listen")
	}

	server := grpc.NewServer()
	sample.RegisterSampleServiceServer(server, sampleServer)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
