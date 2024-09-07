package main

import (
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

type Sensor[T constraints.Integer] interface {
	GenerateData() T
}

type BreathingSensor struct{}

func (s BreathingSensor) GenerateData() int {
	return rand.IntN(100)
}

type SoundSensor struct{}

func (s SoundSensor) GenerateData() int {
	return rand.IntN(200)
}
