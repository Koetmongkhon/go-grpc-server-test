package main

import (
	"log"
	"net"

	"github.com/koetmongkhon/go-grpc/tutorial"
	"google.golang.org/grpc"
)

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":54321")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	tutorial.RegisterTutorialServer(s, &tutorial.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
