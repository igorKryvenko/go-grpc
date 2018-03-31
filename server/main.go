package main

import (
	"fmt"
	"go-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 777))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := api.Server{}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")

	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
