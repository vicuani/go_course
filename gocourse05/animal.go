package main

import "math/rand/v2"

type animalSpecies string

func (as animalSpecies) String() string {
	return string(as)
}

var allAnimalSpecies = []animalSpecies{"Anaconda", "Bear", "Cobra", "Lion", "Tiger"}

func generateRandomAnimalSpecies() animalSpecies {
	return allAnimalSpecies[rand.IntN(len(allAnimalSpecies))]
}

type animalState string

func (s animalState) String() string {
	return string(s)
}

var allAnimalStates = []animalState{"Sleeping", "Running", "Washing", "Eating", "Fighting", "GettingSick", "Escaping"}
var dangerousAnimalStates = allAnimalStates[4:]

func generateRandomAnimalState() animalState {
	return allAnimalStates[rand.IntN(len(allAnimalStates))]
}

func isAnimalStateDangerous(s animalState) bool {
	for _, ds := range dangerousAnimalStates {
		if s == ds {
			return true
		}
	}
	return false
}

type animal struct {
	id      int
	species animalSpecies
	state   animalState
}

func newAnimal(id int) *animal {
	return &animal{
		id:      id,
		species: generateRandomAnimalSpecies(),
		state:   animalState("Sleeping"),
	}
}

type animalHistory map[int][]*animal
