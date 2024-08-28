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

func randName() string {
	b := make([]rune, nameLength)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func (s *Sector) GenerateAndAddAnimal() {
	animal := &Animal{
		id:         len(s.animals),
		animalType: s.area.animalType,
		name:       randName(),
	}

	s.animals = append(s.animals, animal)
}

func (s *Sector) GetAnimalIndex(an *Animal) int {
	for index, currentAnimal := range s.animals {
		if an == currentAnimal {
			return index
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

func (s *Sector) MoveAnimalToOtherSector(os *Sector, an *Animal) error {
	oldIndex := s.GetAnimalIndex(an)
	if oldIndex == -1 {
		return fmt.Errorf("previous sector doesn't contains such an animal")
	}

	newIndex := s.GetAnimalIndex(an)
	if newIndex != -1 {
		return fmt.Errorf("new sector already contains such an animal")
	}

	os.animals = append(os.animals, an)
	s.animals = append(s.animals[:oldIndex], s.animals[oldIndex+1:]...)
	return nil
}
