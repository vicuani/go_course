package main

import (
	"log"
	"net"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	s := grpc.NewServer()
	grpcapi.RegisterTaxiServiceServer(s, newServer())

	log.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
