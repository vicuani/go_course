package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc"
)

type Review struct {
	CargoState    []int32
	DriverService []int32
	DeliverySpeed []int32
}

type server struct {
	grpcapi.UnimplementedTaxiServiceServer
	reviews map[int32]*Review
	mu      sync.Mutex
}

func (s *server) EvaluateCargoState(ctx context.Context, req *grpcapi.EvaluateCargoStateRequest) (*grpcapi.EvaluateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.reviews[req.DriverID] == nil {
		s.reviews[req.DriverID] = &Review{}
	}
	s.reviews[req.DriverID].CargoState = append(s.reviews[req.DriverID].CargoState, req.CargoState)

	return &grpcapi.EvaluateResponse{Message: "Cargo state rating added"}, nil
}

func (s *server) EvaluateDriverService(ctx context.Context, req *grpcapi.EvaluateDriverServiceRequest) (*grpcapi.EvaluateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.reviews[req.DriverID] == nil {
		s.reviews[req.DriverID] = &Review{}
	}
	s.reviews[req.DriverID].DriverService = append(s.reviews[req.DriverID].DriverService, req.DriverService)

	return &grpcapi.EvaluateResponse{Message: "Driver service rating added"}, nil
}

func (s *server) EvaluateDeliverySpeed(ctx context.Context, req *grpcapi.EvaluateDeliverySpeedRequest) (*grpcapi.EvaluateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.reviews[req.DriverID] == nil {
		s.reviews[req.DriverID] = &Review{}
	}
	s.reviews[req.DriverID].DeliverySpeed = append(s.reviews[req.DriverID].DeliverySpeed, req.DeliverySpeed)

	return &grpcapi.EvaluateResponse{Message: "Delivery speed rating added"}, nil
}

func (s *server) DriverReviewsHistory(ctx context.Context, req *grpcapi.DriverReviewsHistoryRequest) (*grpcapi.DriverReviewsHistoryResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	review, ok := s.reviews[req.DriverID]
	if !ok {
		return &grpcapi.DriverReviewsHistoryResponse{
			CargoState: []int32{},
			DriverService: []int32{},
			DeliverySpeed: []int32{},
		}, nil
	}

	return &grpcapi.DriverReviewsHistoryResponse{
		CargoState: review.CargoState,
		DriverService: review.DriverService,
		DeliverySpeed: review.DeliverySpeed,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	s := grpc.NewServer()
	grpcapi.RegisterTaxiServiceServer(s, &server{reviews: make(map[int32]*Review)})

	log.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
