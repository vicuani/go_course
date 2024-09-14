package main

import (
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

type Sensor[T constraints.Integer] interface {
	Measure() T
}

type BreathingSensor struct{}

func (s BreathingSensor) Measure() int {
	return rand.IntN(100)
}

type SoundSensor struct{}

func (s SoundSensor) Measure() int {
	return rand.IntN(200)
}
