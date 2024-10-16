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

func (t *TaxiServiceClient) sendCargoStateRating(driverID int32, cargoState grpcapi.Rating) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateCargoStateRequest{
		DriverId:   driverID,
		CargoState: cargoState,
	}

	res, err := t.client.EvaluateCargoState(ctx, req)
	if err != nil {
		log.Panicf("could not submit cargo state rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) sendDriverServiceRating(driverID int32, driverService grpcapi.Rating) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDriverServiceRequest{
		DriverId:      driverID,
		DriverService: driverService,
	}

	res, err := t.client.EvaluateDriverService(ctx, req)
	if err != nil {
		log.Panicf("could not submit driver service rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) sendDeliverySpeedRating(driverID int32, deliverySpeed grpcapi.Rating) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDeliverySpeedRequest{
		DriverId:      driverID,
		DeliverySpeed: deliverySpeed,
	}

	res, err := t.client.EvaluateDeliverySpeed(ctx, req)
	if err != nil {
		log.Panicf("could not submit delivery speed rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func (t *TaxiServiceClient) getDriverReviews(ctx context.Context, driverID int32) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req := &grpcapi.DriverReviewsHistoryRequest{
		DriverId: driverID,
	}

	res, err := t.client.DriverReviewsHistory(ctx, req)
	if err != nil {
		log.Panicf("could not get driver reviews: %v", err)
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
	taxiClient := TaxiServiceClient{client: client}

	taxiClient.sendCargoStateRating(driverID, grpcapi.Rating_RATING_EXCELLENT)
	taxiClient.sendCargoStateRating(driverID, grpcapi.Rating_RATING_GREAT)
	taxiClient.sendCargoStateRating(driverID, grpcapi.Rating_RATING_EXCELLENT)
	taxiClient.sendCargoStateRating(driverID, grpcapi.Rating_RATING_GOOD)
	taxiClient.sendCargoStateRating(driverID, grpcapi.Rating_RATING_EXCELLENT)

	taxiClient.sendDriverServiceRating(driverID, grpcapi.Rating_RATING_GOOD)
	taxiClient.sendDriverServiceRating(driverID, grpcapi.Rating_RATING_GREAT)

	taxiClient.sendDeliverySpeedRating(driverID, grpcapi.Rating_RATING_EXCELLENT)
	taxiClient.sendDeliverySpeedRating(driverID, grpcapi.Rating_RATING_EXCELLENT)
	taxiClient.sendDeliverySpeedRating(driverID, grpcapi.Rating_RATING_GREAT)

	taxiClient.getDriverReviews(driverID)
}
