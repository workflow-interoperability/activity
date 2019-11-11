package main

import (
	"fmt"
	"log"
	"net"

	mygrpc "github.com/workflow-interoperability/activity/grpc"
	"github.com/workflow-interoperability/activity/parent"
	"github.com/workflow-interoperability/activity/sub"

	"google.golang.org/grpc"
)

func main() {
	go func() {
		var grpcPort = "8083"
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
		if err != nil {
			log.Fatalf("failed to listen grpc: %v", err)
		}
		log.Printf("Listening on: %s", grpcPort)
		gs := grpc.NewServer()
		mygrpc.RegisterActivityServer(gs, &sub.SubActivity{})
		gs.Serve(lis)
	}()
	var grpcPort = "8082"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	log.Printf("Listening on: %s", grpcPort)
	gs := grpc.NewServer()
	mygrpc.RegisterActivityServer(gs, &parent.ParentActivity{})
	gs.Serve(lis)
}
