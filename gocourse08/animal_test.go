package main

import (
	"testing"
)

func TestDetermineAnimalType(t *testing.T) {
	testCases := []struct {
		name        string
		pulse       int
		temperature float64
		expected    string
	}{
		{
			name:        "Bear when pulse <= 60",
			pulse:       60,
			temperature: 30,
			expected:    "Bear",
		},
		{
			name:        "Gorilla when pulse <= 90 and temperature > 37.5",
			pulse:       90,
			temperature: 37.6,
			expected:    "Gorilla",
		},
		{
			name:        "Ape when 60 < pulse <= 90 and temperature <= 37.5",
			pulse:       90,
			temperature: 37.5,
			expected:    "Ape",
		},
		{
			name:        "Lion when pulse > 90 and temperature > 38.5",
			pulse:       91,
			temperature: 38.6,
			expected:    "Lion",
		},
		{
			name:        "Tiger when pulse > 90 and temperature <= 38.5",
			pulse:       91,
			temperature: 38.5,
			expected:    "Tiger",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := determineAnimalType(tt.pulse, tt.temperature)
			if result != tt.expected {
				t.Errorf("determineAnimalType(%d, %.1f) = %s; expected %s", tt.pulse, tt.temperature, result, tt.expected)
			}
		})
	}
}
