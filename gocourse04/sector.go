package main

import (
	"fmt"
	"math/rand/v2"
)

type Sector struct {
	animals      []*Animal
	utilitySpace *UtilitySpace
	area         *Area
}

func NewSector(a *Area) *Sector {
	sector := &Sector{
		animals: []*Animal{},
		area:    a,
	}
	sector.utilitySpace = NewUtilitySpace(sector)
	return sector
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const nameLength int = 10

func (s *Sector) AddAnimal(an *Animal) {
	s.animals = append(s.animals, an)
}

func (s *Sector) NewAnimal(name string) *Animal {
	return NewAnimal(len(s.animals), s.area.animalType, name)
}

func (s *Sector) NewRandomAnimal() *Animal {
	return s.NewAnimal(randName())
}

func randName() string {
	b := make([]rune, nameLength)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func (s *Sector) ContainsAnimal(an *Animal) bool {
	return s.AnimalIndex(an) != -1
}

func (s *Sector) AnimalIndex(an *Animal) int {
	for i, currentAnimal := range s.animals {
		if an == currentAnimal {
			return i
		}
	}
	return -1
}

func (s *Sector) FindAnimalByName(name string) *Animal {
	for _, animal := range s.animals {
		if animal.name == name {
			return animal
		}
	}
	return nil
}

func (s *Sector) MoveAnimalToOtherSector(newSector *Sector, an *Animal) error {
	oldIndex := s.AnimalIndex(an)
	if oldIndex == -1 {
		return fmt.Errorf("previous sector doesn't contains such an animal")
	}

	newIndex := newSector.AnimalIndex(an)
	if newIndex != -1 {
		return fmt.Errorf("new sector already contains such an animal")
	}

	newSector.animals = append(newSector.animals, an)
	s.animals = append(s.animals[:oldIndex], s.animals[oldIndex+1:]...)

	return nil
}
