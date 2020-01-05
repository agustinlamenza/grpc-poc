package main

import (
	"grpc-poc/api"
	"grpc-poc/server/handlers"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}

	s := grpc.NewServer()
	srv := &handlers.Server{}
	api.RegisterCalculatorServiceServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error! : %v", err)
	}
}
