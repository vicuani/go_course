package main

import (
	"context"
	"log"
	"os"
	"time"
)

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