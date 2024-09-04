package main

import "time"

type AnimalData struct {
	animalType  string
	pulse       int
	temperature float64
	breathing   []int
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
	if pulse <= 60 {
		return "Bear"
	} else if pulse > 60 && pulse <= 90 {
		if temperature > 37.5 {
			return "Gorilla"
		} else {
			return "Ape"
		}
	} else if temperature > 38.5 {
		return "Lion"
	} else {
		return "Tiger"
	}
}
