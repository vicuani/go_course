// usage: firstly run postgres in docker:
// docker run --name postgres-container -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

package main

import (
	"log"
	"os"
)

func main() {
	fileUsers, err := os.Open("users.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileUsers.Close()
	users, err := readUsers(fileUsers)
	if err != nil {
		log.Fatal(err)
	}

	fileStatistics, err := os.Open("statistics.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileStatistics.Close()
	statistics, err := readStatistics(fileStatistics)
	if err != nil {
		log.Fatal(err)
	}

	err = write(users, statistics)
	if err != nil {
		log.Fatal(err)
	}
	read()
}
