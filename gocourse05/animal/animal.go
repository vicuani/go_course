package animal

import (
	"math/rand/v2"
)

type animalSpecies string

func (as animalSpecies) String() string {
	return string(as)
}

var allAnimalSpecies = []animalSpecies{"Anaconda", "Bear", "Cobra", "Lion", "Tiger"}

func generateRandomAnimalSpecies() animalSpecies {
	return allAnimalSpecies[rand.IntN(len(allAnimalSpecies))]
}

type AnimalState string

func (s AnimalState) String() string {
	return string(s)
}

var (
	allAnimalStates       = []AnimalState{"Sleeping", "Running", "Washing", "Eating", "Fighting", "GettingSick", "Escaping"}
	dangerousAnimalStates = allAnimalStates[4:]
)

type Animal struct {
	id      int
	species animalSpecies
	state   AnimalState
}

func NewAnimal(id int) *Animal {
	return &Animal{
		id:      id,
		species: generateRandomAnimalSpecies(),
		state:   AnimalState("Sleeping"),
	}
}

func (an *Animal) SetRandomState() {
	an.state = allAnimalStates[rand.IntN(len(allAnimalStates))]
}

func (an *Animal) IsAnimalStateDangerous() bool {
	for _, ds := range dangerousAnimalStates {
		if an.state == ds {
			return true
		}
	}
	return false
}
