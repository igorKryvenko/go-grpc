package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"go-grpc/api"

	"google.golang.org/grpc/credentials"
)

func main() {
	var conn * grpc.ClientConn

	creds,err := credentials.NewClientTLSFromFile("../cert/server.crt","")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn,err = grpc.Dial("igor:777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response,err := c.SayHello(context.Background(), &api.PingMessage{Greeeting: "foo"})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeeting)
}
