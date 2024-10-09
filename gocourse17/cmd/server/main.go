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
	CargoState    []grpcapi.Enum
	DriverService []grpcapi.Enum
	DeliverySpeed []grpcapi.Enum
}

type server struct {
	grpcapi.UnimplementedTaxiServiceServer

	reviewsMu sync.Mutex
	reviews   map[int32]*Review
}

func newServer() *server {
	return &server{reviews: make(map[int32]*Review)}
}

func (s *server) EvaluateCargoState(ctx context.Context, req *grpcapi.EvaluateCargoStateRequest) (*grpcapi.EvaluateCargoStateResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].CargoState = append(s.reviews[req.DriverId].CargoState, req.CargoState)

	return &grpcapi.EvaluateCargoStateResponse{Message: "Cargo state rating added"}, nil
}

func (s *server) EvaluateDriverService(ctx context.Context, req *grpcapi.EvaluateDriverServiceRequest) (*grpcapi.EvaluateDriverServiceResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DriverService = append(s.reviews[req.DriverId].DriverService, req.DriverService)

	return &grpcapi.EvaluateDriverServiceResponse{Message: "Driver service rating added"}, nil
}

func (s *server) EvaluateDeliverySpeed(ctx context.Context, req *grpcapi.EvaluateDeliverySpeedRequest) (*grpcapi.EvaluateDeliverySpeedResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DeliverySpeed = append(s.reviews[req.DriverId].DeliverySpeed, req.DeliverySpeed)

	return &grpcapi.EvaluateDeliverySpeedResponse{Message: "Delivery speed rating added"}, nil
}

func (s *server) DriverReviewsHistory(ctx context.Context, req *grpcapi.DriverReviewsHistoryRequest) (*grpcapi.DriverReviewsHistoryResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	review, ok := s.reviews[req.DriverId]
	if !ok {
		return &grpcapi.DriverReviewsHistoryResponse{
			CargoStates:    []grpcapi.Enum{},
			DriverServices: []grpcapi.Enum{},
			DeliverySpeeds: []grpcapi.Enum{},
		}, nil
	}

	return &grpcapi.DriverReviewsHistoryResponse{
		CargoStates:    review.CargoState,
		DriverServices: review.DriverService,
		DeliverySpeeds: review.DeliverySpeed,
	}, nil
}

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
