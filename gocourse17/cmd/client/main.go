package main

import (
	"context"
	"log"
	"time"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TaxiServiceClient struct {
	client grpcapi.TaxiServiceClient
}

func (t *TaxiServiceClient) sendCargoStateRating(driverID int32, cargoState grpcapi.Enum) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateCargoStateRequest{
		DriverId:   driverID,
		CargoState: cargoState,
	}

	res, err := t.client.EvaluateCargoState(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit cargo state rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) sendDriverServiceRating(driverID int32, driverService grpcapi.Enum) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDriverServiceRequest{
		DriverId:      driverID,
		DriverService: driverService,
	}

	res, err := t.client.EvaluateDriverService(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit driver service rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) sendDeliverySpeedRating(driverID int32, deliverySpeed grpcapi.Enum) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDeliverySpeedRequest{
		DriverId:      driverID,
		DeliverySpeed: deliverySpeed,
	}

	res, err := t.client.EvaluateDeliverySpeed(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit delivery speed rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) getDriverReviews(driverID int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.DriverReviewsHistoryRequest{
		DriverId: driverID,
	}

	res, err := t.client.DriverReviewsHistory(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not get driver reviews: %v", err)
	}

	log.Printf("\nCargo states ratings for driver: %v\n", driverID)
	for i, cargoState := range res.CargoStates {
		log.Printf("#%v cargo state: %v\n", i, cargoState)
	}

	log.Printf("\nDriver service ratings for driver: %v\n", driverID)
	for i, driverService := range res.DriverServices {
		log.Printf("#%v driver service: %v\n", i, driverService)
	}

	log.Printf("\nDelivery speed ratings for driver: %v\n", driverID)
	for i, deliverySpeed := range res.DeliverySpeeds {
		log.Printf("#%v delivery speed: %v\n", i, deliverySpeed)
	}
	log.Println()
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcapi.NewTaxiServiceClient(conn)
	driverID := int32(567)
	serviceClient := TaxiServiceClient{client: client}

	serviceClient.sendCargoStateRating(driverID, grpcapi.Enum_Excellent)
	serviceClient.sendCargoStateRating(driverID, grpcapi.Enum_Great)
	serviceClient.sendCargoStateRating(driverID, grpcapi.Enum_Excellent)
	serviceClient.sendCargoStateRating(driverID, grpcapi.Enum_Good)
	serviceClient.sendCargoStateRating(driverID, grpcapi.Enum_Excellent)

	serviceClient.sendDriverServiceRating(driverID, grpcapi.Enum_Good)
	serviceClient.sendDriverServiceRating(driverID, grpcapi.Enum_Great)

	serviceClient.sendDeliverySpeedRating(driverID, grpcapi.Enum_Excellent)
	serviceClient.sendDeliverySpeedRating(driverID, grpcapi.Enum_Excellent)
	serviceClient.sendDeliverySpeedRating(driverID, grpcapi.Enum_Great)

	serviceClient.getDriverReviews(driverID)
}
