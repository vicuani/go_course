package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable")
	require.NoError(t, err)

	truncateTables(db)
	defer db.Close()
	defer truncateTables(db)

	users := []User{
		{FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com", Age: 28, Gender: "Female", City: "Lviv", TripsCount: 15, Profession: "Designer"},
		{FirstName: "Bob", LastName: "Brown", Email: "bob@example.com", Age: 34, Gender: "Male", City: "Kharkiv", TripsCount: 20, Profession: "Developer"},
	}
	err = insertUsers(db, users)
	require.NoError(t, err)

	dsn := "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable"
	dbRead, err := connectToDatabase(dsn)
	require.NoError(t, err)

	result, err := getUsers(dbRead)
	require.NoError(t, err)
	assert.Equal(t, len(users), len(result))
	assert.Equal(t, users[0].Email, result[0].Email)
	assert.Equal(t, users[1].Email, result[1].Email)
}

func TestGetStatistics(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable")
	require.NoError(t, err)

	truncateTables(db)
	defer db.Close()
	defer truncateTables(db)

	statistics := []Statistics{
		{City: "Lviv", AgeRange: "18-25", AverageTrips: 5},
		{City: "Kharkiv", AgeRange: "26-35", AverageTrips: 10},
	}
	err = insertStatistics(db, statistics)
	require.NoError(t, err)

	dsn := "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable"
	dbRead, err := connectToDatabase(dsn)
	require.NoError(t, err)

	result, err := getStatistics(dbRead)
	require.NoError(t, err)
	assert.Equal(t, len(statistics), len(result))
	assert.Equal(t, statistics[0].City, result[0].City)
	assert.Equal(t, statistics[1].AgeRange, result[1].AgeRange)
}
