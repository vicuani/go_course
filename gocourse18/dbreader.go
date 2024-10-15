package main

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func getUsers(ctx context.Context, db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}

func getStatistics(ctx context.Context, db *gorm.DB) ([]Statistics, error) {
	var statistics []Statistics
	if err := db.WithContext(ctx).Find(&statistics).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve statistics: %w", err)
	}
	return statistics, nil
}

func displayUsers(users []User) {
	fmt.Println("\nUsers from DB:")
	for _, user := range users {
		fmt.Printf("User: %s %s, Email: %s, City: %s\n", user.FirstName, user.LastName, user.Email, user.City)
	}
}

func displayStatistics(statistics []Statistics) {
	fmt.Println("\nStatistics from DB:")
	for _, stat := range statistics {
		fmt.Printf("City: %s, Age Range: %s, Average Trips: %d\n", stat.City, stat.AgeRange, stat.AverageTrips)
	}
}

func read(ctx context.Context) error {
	dsn := "host=localhost port=5432 user=user password=password dbname=taxi sslmode=disable"
	db, err := connectToDatabase(dsn)
	if err != nil {
		return err
	}

	users, err := getUsers(ctx, db)
	if err != nil {
		return err
	}
	displayUsers(users)

	statistics, err := getStatistics(ctx, db)
	if err != nil {
		return err
	}
	displayStatistics(statistics)

	return nil
}