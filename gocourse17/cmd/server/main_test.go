package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func startTestServer() *grpc.Server {
	s := grpc.NewServer()
	srv := server{reviews: make(map[int32]*Review)}
	grpcapi.RegisterTaxiServiceServer(s, &srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return s
}

func TestEvaluateCargoState(t *testing.T) {
	srv := startTestServer()
	defer srv.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Deprecated method DialContext is used specially for mock-server
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Failed to dial: %v", err)
    }
	defer conn.Close()

	client := grpcapi.NewTaxiServiceClient(conn)

	req := &grpcapi.EvaluateCargoStateRequest{
		DriverId:   123,
		CargoState: 5,
	}

	res, err := client.EvaluateCargoState(ctx, req)
	if err != nil {
		t.Fatalf("EvaluateCargoState failed: %v", err)
	}

	if res.Message != "Cargo state rating added" {
		t.Errorf("unexpected result message: %v", res.Message)
	}
}

func TestEvaluateDriverService(t *testing.T) {
	srv := startTestServer()
	defer srv.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Failed to dial: %v", err)
    }
	defer conn.Close()

	client := grpcapi.NewTaxiServiceClient(conn)

	req := &grpcapi.EvaluateDriverServiceRequest{
		DriverId:      234,
		DriverService: 4,
	}

	res, err := client.EvaluateDriverService(ctx, req)
	if err != nil {
		t.Fatalf("EvaluateDriverService failed: %v", err)
	}

	if res.Message != "Driver service rating added" {
		t.Errorf("unexpected result message: %v", res.Message)
	}
}

func TestEvaluateDeliverySpeed(t *testing.T) {
	srv := startTestServer()
	defer srv.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Failed to dial: %v", err)
    }
	defer conn.Close()

	client := grpcapi.NewTaxiServiceClient(conn)

	req := &grpcapi.EvaluateDeliverySpeedRequest{
		DriverId:      135,
		DeliverySpeed: 3,
	}

	res, err := client.EvaluateDeliverySpeed(ctx, req)
	if err != nil {
		t.Fatalf("EvaluateDeliverySpeed failed: %v", err)
	}

	if res.Message != "Delivery speed rating added" {
		t.Errorf("unexpected result message: %v", res.Message)
	}
}

func TestDriverReviewsHistory(t *testing.T) {
	srv := startTestServer()
	defer srv.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Failed to dial: %v", err)
    }
	defer conn.Close()

	client := grpcapi.NewTaxiServiceClient(conn)

	_, err = client.EvaluateCargoState(ctx, &grpcapi.EvaluateCargoStateRequest{DriverId: 246, CargoState: 5})
	if err != nil {
		t.Fatalf("Failed to evaluate cargo state: %v", err)
	}

	_, err = client.EvaluateDriverService(ctx, &grpcapi.EvaluateDriverServiceRequest{DriverId: 246, DriverService: 4})
	if err != nil {
		t.Fatalf("Failed to evaluate driver service: %v", err)
	}

	_, err = client.EvaluateDeliverySpeed(ctx, &grpcapi.EvaluateDeliverySpeedRequest{DriverId: 246, DeliverySpeed: 3})
	if err != nil {
		t.Fatalf("Failed to evaluate delivery speed: %v", err)
	}

	req := &grpcapi.DriverReviewsHistoryRequest{DriverId: 246}
	res, err := client.DriverReviewsHistory(ctx, req)
	if err != nil {
		t.Fatalf("DriverReviewsHistory failed: %v", err)
	}

	if len(res.CargoStates) != 1 || res.CargoStates[0] != 5 {
		t.Errorf("Expected cargo states [5], but got %v", res.CargoStates)
	}
	if len(res.DriverServices) != 1 || res.DriverServices[0] != 4 {
		t.Errorf("Expected driver services [4], but got %v", res.DriverServices)
	}
	if len(res.DeliverySpeeds) != 1 || res.DeliverySpeeds[0] != 3 {
		t.Errorf("Expected delivery speeds [3], but got %v", res.DeliverySpeeds)
	}
}
