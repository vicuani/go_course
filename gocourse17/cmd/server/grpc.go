package main

import (
	"context"
	"sync"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Review struct {
	CargoStates    []grpcapi.Rating
	DriverServices []grpcapi.Rating
	DeliverySpeeds []grpcapi.Rating
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
	if req.DriverId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver ID")
	}

	if req.CargoState == grpcapi.Rating_RATING_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "invalid rating")
	}

	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].CargoStates = append(s.reviews[req.DriverId].CargoStates, req.CargoState)

	return &grpcapi.EvaluateCargoStateResponse{Message: "Cargo state rating added"}, nil
}

func (s *server) EvaluateDriverService(ctx context.Context, req *grpcapi.EvaluateDriverServiceRequest) (*grpcapi.EvaluateDriverServiceResponse, error) {
	if req.DriverId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver ID")
	}

	if req.DriverService == grpcapi.Rating_RATING_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "invalid rating")
	}

	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DriverServices = append(s.reviews[req.DriverId].DriverServices, req.DriverService)

	return &grpcapi.EvaluateDriverServiceResponse{Message: "Driver service rating added"}, nil
}

func (s *server) EvaluateDeliverySpeed(ctx context.Context, req *grpcapi.EvaluateDeliverySpeedRequest) (*grpcapi.EvaluateDeliverySpeedResponse, error) {
	if req.DriverId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver ID")
	}

	if req.DeliverySpeed == grpcapi.Rating_RATING_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "invalid rating")
	}

	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	if s.reviews[req.DriverId] == nil {
		s.reviews[req.DriverId] = &Review{}
	}
	s.reviews[req.DriverId].DeliverySpeeds = append(s.reviews[req.DriverId].DeliverySpeeds, req.DeliverySpeed)

	return &grpcapi.EvaluateDeliverySpeedResponse{Message: "Delivery speed rating added"}, nil
}

func (s *server) DriverReviewsHistory(ctx context.Context, req *grpcapi.DriverReviewsHistoryRequest) (*grpcapi.DriverReviewsHistoryResponse, error) {
	if req.DriverId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver ID")
	}

	s.reviewsMu.Lock()
	defer s.reviewsMu.Unlock()

	review, ok := s.reviews[req.DriverId]
	if !ok {
		return &grpcapi.DriverReviewsHistoryResponse{
			CargoStates:    []grpcapi.Rating{},
			DriverServices: []grpcapi.Rating{},
			DeliverySpeeds: []grpcapi.Rating{},
		}, nil
	}

	return &grpcapi.DriverReviewsHistoryResponse{
		CargoStates:    review.CargoStates,
		DriverServices: review.DriverServices,
		DeliverySpeeds: review.DeliverySpeeds,
	}, nil
}
