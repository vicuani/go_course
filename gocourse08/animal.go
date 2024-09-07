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
		temperature: float64(temperature),
	}
}

func determineAnimalType(pulse int, temperature float64) string {
	switch {
	case pulse <= 60:
		return "Bear"
	case pulse > 60 && pulse <= 90:
		switch {
		case temperature > 37.5:
			return "Gorilla"
		default:
			return "Ape"
		}
	default:
		switch {
		case temperature > 38.5:
			return "Lion"
		default:
			return "Tiger"
		}
	}
}
