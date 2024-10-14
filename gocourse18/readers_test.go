package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadUsers(t *testing.T) {
	data := `first_name,last_name,email,age,gender,city,trips_count,profession
John,Doe,john@example.com,30,Male,Kyiv,5,Driver
Jane,Smith,jane@example.com,25,Female,Odesa,10,Engineer`

	reader := bytes.NewReader([]byte(data))

	users, err := readUsers(reader)
	require.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "John", users[0].FirstName)
	assert.Equal(t, "Doe", users[0].LastName)
}

func TestReadStatistics(t *testing.T) {
	data := `city,age_range,average_trips
Kyiv,18-25,8
Odesa,26-35,12`

	reader := bytes.NewReader([]byte(data))

	statistics, err := readStatistics(reader)
	require.NoError(t, err)
	assert.Equal(t, 2, len(statistics))
	assert.Equal(t, "Kyiv", statistics[0].City)
	assert.Equal(t, "18-25", statistics[0].AgeRange)
}
