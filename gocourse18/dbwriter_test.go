package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const dbtestname = "test_db"

func TestCreateDatabaseIfNotExists(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, err := initDBWriter(ctx, "postgres")
	require.NoError(t, err)
	defer dbw.db.Close()

	err = dbw.createDatabaseIfNotExists(ctx, dbtestname)
	require.NoError(t, err)

	var dbName string
	err = dbw.db.QueryRowContext(ctx, "SELECT datname FROM pg_database WHERE datname = $1", dbtestname).Scan(&dbName)

	require.NoError(t, err)
	assert.Equal(t, dbtestname, dbName)
}

func TestCreateTables(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, err := initDBWriter(ctx, dbtestname)
	require.NoError(t, err)
	defer dbw.db.Close()

	err = dbw.createTables(ctx)
	require.NoError(t, err)

	var count int
	err = dbw.db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'users'")
	require.NoError(t, err)
	assert.Equal(t, 1, count)

	err = dbw.db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'statistics'")
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestInsertUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, err := initDBWriter(ctx, dbtestname)
	require.NoError(t, err)

	defer dbw.db.Close()
	defer dbw.truncateTables(ctx)

	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 30, Gender: "Male", City: "Kyiv", TripsCount: 5, Profession: "Driver"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 25, Gender: "Female", City: "Odesa", TripsCount: 10, Profession: "Engineer"},
	}

	err = dbw.insertUsers(ctx, users)
	require.NoError(t, err)

	var count int
	err = dbw.db.Get(&count, "SELECT COUNT(*) FROM users")
	require.NoError(t, err)
	assert.Equal(t, len(users), count)
}

func TestInsertStatistics(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, err := initDBWriter(ctx, dbtestname)
	require.NoError(t, err)

	defer dbw.db.Close()
	defer dbw.truncateTables(ctx)

	statistics := []Statistics{
		{City: "Kyiv", AgeRange: "18-25", AverageTrips: 8},
		{City: "Odesa", AgeRange: "26-35", AverageTrips: 12},
	}

	err = dbw.insertStatistics(ctx, statistics)
	require.NoError(t, err)

	var count int
	err = dbw.db.Get(&count, "SELECT COUNT(*) FROM statistics")
	require.NoError(t, err)
	assert.Equal(t, len(statistics), count)
}
