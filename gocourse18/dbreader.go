package main

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBReader struct {
	db *gorm.DB
}

func initDBReader(dbconfig string) (*DBReader, error) {
	db, err := gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DBReader{
		db: db,
	}, nil
}

func (dbr *DBReader) getUsers(ctx context.Context) ([]User, error) {
	var users []User
	if err := dbr.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}

func (dbr *DBReader) getStatistics(ctx context.Context) ([]Statistics, error) {
	var statistics []Statistics
	if err := dbr.db.WithContext(ctx).Find(&statistics).Error; err != nil {
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
	dbr, err := initDBReader(createDBConnectionPath("taxi"))
	if err != nil {
		return err
	}

	users, err := dbr.getUsers(ctx)
	if err != nil {
		return err
	}
	displayUsers(users)

	statistics, err := dbr.getStatistics(ctx)
	if err != nil {
		return err
	}
	displayStatistics(statistics)

	return nil
}
