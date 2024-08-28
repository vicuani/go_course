package animal

import (
	"fmt"
	"math/rand/v2"
)

const MAX_INDICATOR_VALUE = 100
const CRITICAL_INDICATOR_VALUE = 20

type Animal struct {
	id     int
	Health int
	Hunger int
	Mood   int
}

func NewAnimal(id int) *Animal {
	return &Animal{
		id:     id,
		Health: MAX_INDICATOR_VALUE,
		Hunger: MAX_INDICATOR_VALUE,
		Mood:   MAX_INDICATOR_VALUE,
	}
}

func (an *Animal) RandomlyChangeIndicators() {
	an.Health = rand.IntN(MAX_INDICATOR_VALUE + 1)
	an.Hunger = rand.IntN(MAX_INDICATOR_VALUE + 1)
	an.Mood = rand.IntN(MAX_INDICATOR_VALUE + 1)

	fmt.Printf("Animal randomly changed it's vlaues: %v, has critical values: %v\n", *an, an.HasCriticalValues())
}

func (an *Animal) HasCriticalValues() bool {
	return an.Health < CRITICAL_INDICATOR_VALUE || an.Hunger < CRITICAL_INDICATOR_VALUE || an.Mood < CRITICAL_INDICATOR_VALUE
}

type Enclosure struct {
	ID       int
	IsOpened bool
}

func NewEnclosure(id int) *Enclosure {
	return &Enclosure{
		ID:       id,
		IsOpened: rand.IntN(2) == 1,
	}
}

type Feeder struct {
	ID      int
	IsEmpty bool
}

func NewFeeder(id int) *Feeder {
	return &Feeder{
		ID:      id,
		IsEmpty: rand.IntN(2) == 1,
	}
}

func GenerateAnimals(n int) []*Animal {
	var animals []*Animal
	for i := 0; i < n; i++ {
		animal := NewAnimal(i)
		animals = append(animals, animal)
	}
	return animals
}

func GenerateEnclosures(n int) []*Enclosure {
	var enclosures []*Enclosure
	for i := 0; i < n; i++ {
		enclosure := NewEnclosure(i)
		enclosures = append(enclosures, enclosure)
	}
	return enclosures
}

func GenerateFeeders(n int) []*Feeder {
	var feeders []*Feeder
	for i := 0; i < n; i++ {
		feeder := NewFeeder(i)
		feeders = append(feeders, feeder)
	}
	return feeders
}
