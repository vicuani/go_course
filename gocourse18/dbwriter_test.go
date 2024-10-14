package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDatabaseIfNotExists(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	err = createDatabaseIfNotExists(db, "test_db")
	require.NoError(t, err)

	var dbName string
	err = db.QueryRow("SELECT datname FROM pg_database WHERE datname = 'test_db'").Scan(&dbName)
	assert.NoError(t, err)
	assert.Equal(t, "test_db", dbName)
}

func TestCreateTables(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	err = createTables(db)
	require.NoError(t, err)

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'users'")
	require.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'statistics'")
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestInsertUsers(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	defer truncateTables(db)

	users := []User{
		{FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 30, Gender: "Male", City: "Kyiv", TripsCount: 5, Profession: "Driver"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 25, Gender: "Female", City: "Odesa", TripsCount: 10, Profession: "Engineer"},
	}

	err = insertUsers(db, users)
	assert.NoError(t, err)

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM users")
	assert.NoError(t, err)
	assert.Equal(t, len(users), count)
}

func TestInsertStatistics(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=test_db sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	defer truncateTables(db)

	statistics := []Statistics{
		{City: "Kyiv", AgeRange: "18-25", AverageTrips: 8},
		{City: "Odesa", AgeRange: "26-35", AverageTrips: 12},
	}

	err = insertStatistics(db, statistics)
	assert.NoError(t, err)

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM statistics")
	assert.NoError(t, err)
	assert.Equal(t, len(statistics), count)
}
