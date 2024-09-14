package main

import (
	"strconv"
	"testing"
)

func TestDetermineAnimalType(t *testing.T) {
	type expectedParams struct {
		pulse       int
		temperature float64
		expected    string
	}

	data := []expectedParams{
		{50, 30.0, "Bear"},    // pulse <= 60
		{70, 38.0, "Gorilla"}, // 60 < pulse <= 90, temperature > 37.5
		{70, 36.0, "Ape"},     // 60 < pulse <= 90, temperature <= 37.5
		{100, 39.0, "Lion"},   // pulse > 90, temperature > 38.5
		{100, 37.0, "Tiger"},  // pulse > 90, temperature <= 38.5
	}

	for i, tt := range data {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := determineAnimalType(tt.pulse, tt.temperature)
			if result != tt.expected {
				t.Errorf("determineAnimalType(%d, %.1f) = %s; expected %s", tt.pulse, tt.temperature, result, tt.expected)
			}
		})
	}
}
