package main

import (
	"context"
	"log"
	"time"

	"github.com/vicuani/go_course/gocourse17/internal/grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendCargoStateRating(client grpcapi.TaxiServiceClient, driverID int32, cargoState int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateCargoStateRequest{
		DriverId:   driverID,
		CargoState: cargoState,
	}

	res, err := client.EvaluateCargoState(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit cargo state rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func sendDriverServiceRating(client grpcapi.TaxiServiceClient, driverID int32, driverService int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDriverServiceRequest{
		DriverId:      driverID,
		DriverService: driverService,
	}

	res, err := client.EvaluateDriverService(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit driver service rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func sendDeliverySpeedRating(client grpcapi.TaxiServiceClient, driverID int32, deliverySpeed int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.EvaluateDeliverySpeedRequest{
		DriverId:      driverID,
		DeliverySpeed: deliverySpeed,
	}

	res, err := client.EvaluateDeliverySpeed(ctx, req)
	if err != nil {
		cancel()
		log.Fatalf("could not submit delivery speed rating: %v", err)
	}

	log.Printf("Response from server: %s\n", res.Message)
}

func getDriverReviews(client grpcapi.TaxiServiceClient, driverID int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcapi.DriverReviewsHistoryRequest{
		DriverId: driverID,
	}

	res, err := client.DriverReviewsHistory(ctx, req)
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

	sendCargoStateRating(client, driverID, 5)
	sendCargoStateRating(client, driverID, 4)
	sendCargoStateRating(client, driverID, 5)
	sendCargoStateRating(client, driverID, 3)
	sendCargoStateRating(client, driverID, 5)

	sendDriverServiceRating(client, driverID, 3)
	sendDriverServiceRating(client, driverID, 4)

	sendDeliverySpeedRating(client, driverID, 5)
	sendDeliverySpeedRating(client, driverID, 5)
	sendDeliverySpeedRating(client, driverID, 4)

	getDriverReviews(client, driverID)
}
