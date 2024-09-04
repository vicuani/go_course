package main

import "math/rand/v2"

type Sensor[T any] interface {
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
