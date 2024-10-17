package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadUsers(t *testing.T) {
	t.Run("correct data", func(t *testing.T) {
		data := `first_name,last_name,email,age,gender,city,trips_count,profession
John,Doe,john@example.com,30,Male,Kyiv,5,Driver
Jane,Smith,jane@example.com,25,Female,Odesa,10,Engineer`

		reader := strings.NewReader(data)

		users, err := readUsers(reader)
		require.NoError(t, err)
		assert.Equal(t, []User{
			{
				ID:         1,
				FirstName:  "John",
				LastName:   "Doe",
				Email:      "john@example.com",
				Age:        30,
				Gender:     "Male",
				City:       "Kyiv",
				TripsCount: 5,
				Profession: "Driver",
			},
			{
				ID:         2,
				FirstName:  "Jane",
				LastName:   "Smith",
				Email:      "jane@example.com",
				Age:        25,
				Gender:     "Female",
				City:       "Odesa",
				TripsCount: 10,
				Profession: "Engineer",
			},
		}, users)
	})

	t.Run("incorrect data", func(t *testing.T) {
		invalidData := `first_name,last_name,email,age,gender,city,trips_count,profession
Oleh,Liashko,radykal@tut.com,thirty,Male,Kyiv,seven,Driver`

		invalidReader := strings.NewReader(invalidData)

		_, err := readUsers(invalidReader)
		require.Error(t, err)
	})
}

func TestReadStatistics(t *testing.T) {
	t.Run("correct data", func(t *testing.T) {
		data := `city,age_range,average_trips
Kyiv,18-25,8
Odesa,26-35,12`

		reader := strings.NewReader(data)

		statistics, err := readStatistics(reader)
		require.NoError(t, err)

		assert.Equal(t, []Statistics{
			{
				City:         "Kyiv",
				AgeRange:     "18-25",
				AverageTrips: 8,
			},
			{
				City:         "Odesa",
				AgeRange:     "26-35",
				AverageTrips: 12,
			},
		}, statistics)
	})

	t.Run("incorrect data", func(t *testing.T) {
		invalidData := `city,age_range,average_trips
Kyiv,18-25,eight`

		invalidReader := strings.NewReader(invalidData)

		_, err := readStatistics(invalidReader)
		require.Error(t, err)
	})
}
