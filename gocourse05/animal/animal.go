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

type animalState string

func (s animalState) String() string {
	return string(s)
}

var allAnimalStates = []animalState{"Sleeping", "Running", "Washing", "Eating", "Fighting", "GettingSick", "Escaping"}
var dangerousAnimalStates = allAnimalStates[4:]

func GenerateRandomAnimalState() animalState {
	return allAnimalStates[rand.IntN(len(allAnimalStates))]
}

func IsAnimalStateDangerous(s animalState) bool {
	for _, ds := range dangerousAnimalStates {
		if s == ds {
			return true
		}
	}
	return false
}

type Animal struct {
	ID      int
	species animalSpecies
	State   animalState
}

func NewAnimal(id int) *Animal {
	return &Animal{
		ID:      id,
		species: generateRandomAnimalSpecies(),
		State:   animalState("Sleeping"),
	}
}

// Animals themselves are adding to history, because their state can be changed
type Animals []Animal

type HistoryEpisode interface {
	Add(an *Animal)
	GetData() Animals
}

type FullHistoryEpisode struct {
	data Animals
}

func CreateFullHistoryEpisode() HistoryEpisode {
	return &FullHistoryEpisode{
		data: Animals{},
	}
}

func (fh *FullHistoryEpisode) Add(an *Animal) {
	fh.data = append(fh.data, *an)
}

func (fh *FullHistoryEpisode) GetData() Animals {
	return fh.data
}

type DangerousHistoryEpisode struct {
	data Animals
}

func CreateDangerousHistoryEpisode() HistoryEpisode {
	return &DangerousHistoryEpisode{
		data: Animals{},
	}
}

func (dh *DangerousHistoryEpisode) Add(an *Animal) {
	if !IsAnimalStateDangerous(an.State) {
		return
	}
	dh.data = append(dh.data, *an)
}

func (dh *DangerousHistoryEpisode) GetData() Animals {
	return dh.data
}
