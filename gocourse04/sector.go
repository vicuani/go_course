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

func (s *Sector) CreateAndAddAnimal(id int, name string) {
	s.animals = append(s.animals, NewAnimal(id, s.area.animalType, name))
}

func (s *Sector) GenerateAndAddAnimal() {
	s.CreateAndAddAnimal(len(s.animals), randName())
}

func randName() string {
	b := make([]rune, nameLength)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
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

func (s *Sector) MoveAnimalToOtherSector(newSector *Sector, an *Animal) error {
	oldIndex := s.GetAnimalIndex(an)
	if oldIndex == -1 {
		return fmt.Errorf("previous sector doesn't contains such an animal")
	}

	newIndex := newSector.GetAnimalIndex(an)
	if newIndex != -1 {
		return fmt.Errorf("new sector already contains such an animal")
	}

	newSector.animals = append(newSector.animals, an)
	s.animals = append(s.animals[:oldIndex], s.animals[oldIndex+1:]...)
	newIndex = len(newSector.animals) - 1

	fmt.Printf("Animal was moved from one sector to other: old index = %v, new index = %v \n", oldIndex, newIndex)
	return nil
}
