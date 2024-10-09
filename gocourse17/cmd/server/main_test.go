package main

import (
	"context"
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

func startTestServer(t *testing.T) (dialer func(context.Context, string) (net.Conn, error)) {
	lis := bufconn.Listen(bufSize)

	baseServer := grpc.NewServer()
	grpcapi.RegisterTaxiServiceServer(baseServer, newServer())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			t.Errorf("error serving server: %v", err)
		}
	}()

	t.Cleanup(func() {
		err := lis.Close()
		if err != nil {
			t.Fatalf("error closing listener: %v", err)
		}
		baseServer.Stop()
	})

	return func(ctx context.Context, _ string) (net.Conn, error) { return lis.DialContext(ctx) }
}

func newTestClient(t *testing.T, ctx context.Context, dialer func(context.Context, string) (net.Conn, error)) grpcapi.TaxiServiceClient {
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("error connecting to server: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(); err != nil {
			t.Fatalf("error closing connection: %v", err)
		}
	})

	return grpcapi.NewTaxiServiceClient(conn)
}

func TestEvaluateCargoState(t *testing.T) {
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

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
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

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
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

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
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

	_, err := client.EvaluateCargoState(ctx, &grpcapi.EvaluateCargoStateRequest{DriverId: 246, CargoState: 5})
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
