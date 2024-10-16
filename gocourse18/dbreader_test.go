package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(ctx context.Context, t *testing.T) (*DBWriter, *DBReader) {
	dbw, err := initDBWriter(ctx, dbtestname)
	require.NoError(t, err)

	dbw.createTables(ctx)
	dbw.truncateTables(ctx)

	t.Cleanup(func() {
		dbw.truncateTables(ctx)
		dbw.db.Close()
	})

	dbr, err := initDBReader(createDBConnectionPath(dbtestname))
	require.NoError(t, err)

	return dbw, dbr
}

func TestGetUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, dbr := setup(ctx, t)

	users := []User{
		{ID: 1, FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com", Age: 28, Gender: "Female", City: "Lviv", TripsCount: 15, Profession: "Designer"},
		{ID: 2, FirstName: "Bob", LastName: "Brown", Email: "bob@example.com", Age: 34, Gender: "Male", City: "Kharkiv", TripsCount: 20, Profession: "Developer"},
	}
	err := dbw.insertUsers(ctx, users)
	require.NoError(t, err)

	result, err := dbr.getUsers(ctx)
	require.NoError(t, err)

	assert.Equal(t, users, result)
}

func TestGetStatistics(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbw, dbr := setup(ctx, t)

	statistics := []Statistics{
		{City: "Lviv", AgeRange: "18-25", AverageTrips: 5},
		{City: "Kharkiv", AgeRange: "26-35", AverageTrips: 10},
	}
	err := dbw.insertStatistics(ctx, statistics)
	require.NoError(t, err)

	result, err := dbr.getStatistics(ctx)
	require.NoError(t, err)

	assert.Equal(t, statistics, result)
}
