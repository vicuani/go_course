package main

import "time"

type AnimalData struct {
	animalType  string
	pulse       int
	temperature float64
	breaths     []int
	sounds      []int
	timestamp   time.Time
}

func NewAnimalData(pulse int, temperature float64) *AnimalData {
	return &AnimalData{
		animalType:  determineAnimalType(pulse, temperature),
		pulse:       pulse,
		temperature: temperature,
	}
}

func determineAnimalType(pulse int, temperature float64) string {
	switch {
	case pulse <= 60:
		return "Bear"
	case pulse <= 90:
		if temperature > 37.5 {
			return "Gorilla"
		}
		return "Ape"
	default:
		if temperature > 38.5 {
			return "Lion"
		}
		return "Tiger"
	}
}
