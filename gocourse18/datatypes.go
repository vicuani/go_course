package main

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
