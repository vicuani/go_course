package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func createDBConnectionPath(dbName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), dbName)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fileUsers, err := os.Open("users.csv")
	if err != nil {
		log.Panic(err)
	}
	defer fileUsers.Close()

	users, err := readUsers(fileUsers)
	if err != nil {
		log.Panic(err)
	}

	fileStatistics, err := os.Open("statistics.csv")
	if err != nil {
		log.Panic(err)
	}
	defer fileStatistics.Close()

	statistics, err := readStatistics(fileStatistics)
	if err != nil {
		log.Panic(err)
	}

	err = write(ctx, users, statistics)
	if err != nil {
		log.Panic(err)
	}

	err = read(ctx)
	if err != nil {
		log.Panic(err)
	}
}
