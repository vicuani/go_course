package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID         int    `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	Age        int    `db:"age"`
	Gender     string `db:"gender"`
	City       string `db:"city"`
	TripsCount int    `db:"trips_count"`
	Profession string `db:"profession"`
}

type Statistics struct {
	City         string `db:"city"`
	AgeRange     string `db:"age_range"`
	AverageTrips int    `db:"average_trips"`
}

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE,
    age INT,
    gender TEXT,
    city TEXT,
    trips_count INT,
    profession TEXT
);
CREATE TABLE IF NOT EXISTS statistics (
    city TEXT,
    age_range TEXT,
    average_trips DECIMAL,
    PRIMARY KEY (city, age_range)
);`

func readUsers(filename string) ([]User, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []User
	for i, record := range records {
		if i == 0 {
			continue
		}

		age, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		tripsCount, err := strconv.Atoi(record[6])
		if err != nil {
			return nil, err
		}

		user := User{
			FirstName:  record[0],
			LastName:   record[1],
			Email:      record[2],
			Age:        age,
			Gender:     record[4],
			City:       record[5],
			TripsCount: tripsCount,
			Profession: record[7],
		}
		users = append(users, user)
	}
	return users, nil
}

func readStatistics(filename string) ([]Statistics, error) {
	var stats []Statistics

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] {
		averageTrips, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		stats = append(stats, Statistics{
			City:         record[0],
			AgeRange:     record[1],
			AverageTrips: averageTrips,
		})
	}

	return stats, nil
}

func main() {
	users, err := readUsers("users.csv")
	if err != nil {
		panic(err)
	}

	statistics, err := readStatistics("statistics.csv")
	if err != nil {
		panic(err)
	}

	write(users, statistics)
	read()
}

func write(users []User, statistics []Statistics) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	var dbName string
	err = db.QueryRow("SELECT datname FROM pg_database WHERE datname = 'taxi'").Scan(&dbName)

	if err == sql.ErrNoRows {
		_, err = db.Exec("CREATE DATABASE taxi")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB 'taxi' created.")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB 'taxi' already exists.")
	}

	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=taxi sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	for _, user := range users {
		_, err := db.NamedExec(`INSERT INTO users (first_name, last_name, email, age, gender, city, trips_count, profession)
								VALUES (:first_name, :last_name, :email, :age, :gender, :city, :trips_count, :profession)
								ON CONFLICT (email) DO NOTHING;`, &user)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Users added successfully")

	for _, stat := range statistics {
		_, err := db.NamedExec(`INSERT INTO statistics (city, age_range, average_trips)
								VALUES (:city, :age_range, :average_trips)
								ON CONFLICT (city, age_range) DO NOTHING;`, &stat)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Statistics added successfully")
}

func read() {
	dsn := "host=localhost user=user password=password dbname=taxi sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("\nUsers from DB:")
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	for _, user := range users {
		fmt.Printf("User: %s %s, Email: %s, City: %s\n", user.FirstName, user.LastName, user.Email, user.City)
	}

	fmt.Println("\nStatistics from DB:")
	var statistics []Statistics
	if err := db.Find(&statistics).Error; err != nil {
		log.Fatal("Failed to retrieve statistics:", err)
	}

	for _, stat := range statistics {
		fmt.Printf("City: %s, Age Range: %s, Average Trips: %d\n", stat.City, stat.AgeRange, stat.AverageTrips)
	}
}