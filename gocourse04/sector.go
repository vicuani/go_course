package main

import (
	"fmt"
	"math/rand"
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
		b[i] = letters[rand.Intn(len(letters))]
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

func (s *Sector) FindAnimalByName(name string) *Animal {
	for _, animal := range s.animals {
		if animal.name == name {
			return animal
		}
	}
	return nil
}

type UtilitySpace struct {
	sector *Sector
}

func NewUtilitySpace(sector *Sector) *UtilitySpace {
	return &UtilitySpace{
		sector: sector,
	}
}

func (us *UtilitySpace) Cleaning() {
	fmt.Println("Fröken Bock is on hers way")
}

func (us *UtilitySpace) Feeding(animalID int) error {
	//	Find this animal
	animals := us.sector.animals
	if animalID < 0 || animalID >= len(animals) {
		return fmt.Errorf("Animal with such ID doesn't exist")
	}

	fmt.Printf("Feeding animal with id = %v\n", animalID)
	return nil
}
