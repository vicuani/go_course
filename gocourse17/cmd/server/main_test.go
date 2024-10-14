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

	t.Run("incorrect driver id", func(t *testing.T) {
		req := &grpcapi.EvaluateCargoStateRequest{
			DriverId:   -4,
			CargoState: grpcapi.Rating_RATING_EXCELLENT,
		}

		_, err := client.EvaluateCargoState(ctx, req)
		if err == nil {
			t.Fatalf("Expected error 'incorrect driver id' but got no error")
		}
	})

	t.Run("normal case", func(t *testing.T) {
		req := &grpcapi.EvaluateCargoStateRequest{
			DriverId:   123,
			CargoState: grpcapi.Rating_RATING_EXCELLENT,
		}

		res, err := client.EvaluateCargoState(ctx, req)
		if err != nil {
			t.Fatalf("EvaluateCargoState failed: %v", err)
		}

		const expectedMessage = "Cargo state rating added"
		if res.Message != expectedMessage {
			t.Errorf("Expected message: %v, got message: %v", expectedMessage, res.Message)
		}
	})
}

func TestEvaluateDriverService(t *testing.T) {
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

	t.Run("incorrect driver id", func(t *testing.T) {
		req := &grpcapi.EvaluateDriverServiceRequest{
			DriverId:      -1,
			DriverService: grpcapi.Rating_RATING_POOR,
		}

		_, err := client.EvaluateDriverService(ctx, req)
		if err == nil {
			t.Fatalf("Expected error 'incorrect driver id' but got no error")
		}
	})

	t.Run("normal case", func(t *testing.T) {
		req := &grpcapi.EvaluateDriverServiceRequest{
			DriverId:      234,
			DriverService: grpcapi.Rating_RATING_POOR,
		}

		res, err := client.EvaluateDriverService(ctx, req)
		if err != nil {
			t.Fatalf("EvaluateDriverService failed: %v", err)
		}

		const expectedMessage = "Driver service rating added"
		if res.Message != expectedMessage {
			t.Errorf("Expected message: %v, got message: %v", expectedMessage, res.Message)
		}
	})
}

func TestEvaluateDeliverySpeed(t *testing.T) {
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

	t.Run("incorrect driver id", func(t *testing.T) {
		req := &grpcapi.EvaluateDeliverySpeedRequest{
			DriverId:      -6,
			DeliverySpeed: grpcapi.Rating_RATING_FAIR,
		}

		_, err := client.EvaluateDeliverySpeed(ctx, req)
		if err == nil {
			t.Fatalf("Expected error 'incorrect driver id' but got no error")
		}
	})

	t.Run("normal case", func(t *testing.T) {
		req := &grpcapi.EvaluateDeliverySpeedRequest{
			DriverId:      135,
			DeliverySpeed: grpcapi.Rating_RATING_FAIR,
		}

		res, err := client.EvaluateDeliverySpeed(ctx, req)
		if err != nil {
			t.Fatalf("EvaluateDeliverySpeed failed: %v", err)
		}

		const expectedMessage = "Delivery speed rating added"
		if res.Message != expectedMessage {
			t.Errorf("Expected message: %v, got message: %v", expectedMessage, res.Message)
		}
	})
}

func TestDriverReviewsHistory(t *testing.T) {
	dialer := startTestServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := newTestClient(t, ctx, dialer)

	const driverID = 246

	_, err := client.EvaluateCargoState(ctx, &grpcapi.EvaluateCargoStateRequest{DriverId: driverID, CargoState: grpcapi.Rating_RATING_FAIR})
	if err != nil {
		t.Fatalf("Failed to evaluate cargo state: %v", err)
	}

	_, err = client.EvaluateDriverService(ctx, &grpcapi.EvaluateDriverServiceRequest{DriverId: driverID, DriverService: grpcapi.Rating_RATING_EXCELLENT})
	if err != nil {
		t.Fatalf("Failed to evaluate driver service: %v", err)
	}

	_, err = client.EvaluateDeliverySpeed(ctx, &grpcapi.EvaluateDeliverySpeedRequest{DriverId: driverID, DeliverySpeed: grpcapi.Rating_RATING_GOOD})
	if err != nil {
		t.Fatalf("Failed to evaluate delivery speed: %v", err)
	}

	req := &grpcapi.DriverReviewsHistoryRequest{DriverId: driverID}
	res, err := client.DriverReviewsHistory(ctx, req)
	if err != nil {
		t.Fatalf("DriverReviewsHistory failed: %v", err)
	}

	if len(res.CargoStates) != 1 {
		t.Fatalf("Expected to receive one cargo state, but got: %v", len(res.CargoStates))
	}

	if len(res.DriverServices) != 1 {
		t.Fatalf("Expected to receive one driver service, but got: %v", len(res.DriverServices))
	}

	if len(res.DeliverySpeeds) != 1 {
		t.Fatalf("Expected to receive one delivery speed, but got: %v", len(res.DeliverySpeeds))
	}

	if res.CargoStates[0] != grpcapi.Rating_RATING_FAIR {
		t.Errorf("Expected cargo states [Enum_Fair], but got %v", res.CargoStates)
	}
	if res.DriverServices[0] != grpcapi.Rating_RATING_EXCELLENT {
		t.Errorf("Expected driver services [Enum_Excellent], but got %v", res.DriverServices)
	}
	if res.DeliverySpeeds[0] != grpcapi.Rating_RATING_GOOD {
		t.Errorf("Expected delivery speeds [Enum_Good], but got %v", res.DeliverySpeeds)
	}
}
