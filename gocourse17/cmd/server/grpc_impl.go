package main

import (
	"context"
	"sync"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Review struct {
	CargoState    []grpcapi.Rating
	DriverService []grpcapi.Rating
	DeliverySpeed []grpcapi.Rating
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

	if req.DriverId < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid driver ID")
	}

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].CargoState = append(s.reviews[req.DriverId].CargoState, req.CargoState)

	return &grpcapi.EvaluateCargoStateResponse{Message: "Cargo state rating added"}, nil
}

func (s *server) EvaluateDriverService(ctx context.Context, req *grpcapi.EvaluateDriverServiceRequest) (*grpcapi.EvaluateDriverServiceResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if req.DriverId < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid driver ID")
	}

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DriverService = append(s.reviews[req.DriverId].DriverService, req.DriverService)

	return &grpcapi.EvaluateDriverServiceResponse{Message: "Driver service rating added"}, nil
}

func (s *server) EvaluateDeliverySpeed(ctx context.Context, req *grpcapi.EvaluateDeliverySpeedRequest) (*grpcapi.EvaluateDeliverySpeedResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if req.DriverId < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid driver ID")
	}

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DeliverySpeed = append(s.reviews[req.DriverId].DeliverySpeed, req.DeliverySpeed)

	return &grpcapi.EvaluateDeliverySpeedResponse{Message: "Delivery speed rating added"}, nil
}

func (s *server) DriverReviewsHistory(ctx context.Context, req *grpcapi.DriverReviewsHistoryRequest) (*grpcapi.DriverReviewsHistoryResponse, error) {
	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if req.DriverId < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid driver ID")
	}

	review, ok := s.reviews[req.DriverId]
	if !ok {
		return &grpcapi.DriverReviewsHistoryResponse{
			CargoStates:    []grpcapi.Rating{},
			DriverServices: []grpcapi.Rating{},
			DeliverySpeeds: []grpcapi.Rating{},
		}, nil
	}

	return &grpcapi.DriverReviewsHistoryResponse{
		CargoStates:    review.CargoState,
		DriverServices: review.DriverService,
		DeliverySpeeds: review.DeliverySpeed,
	}, nil
}
