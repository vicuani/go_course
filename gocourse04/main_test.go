package main

import "testing"

func TestAnimalGeneration(t *testing.T) {
	area := NewArea("fishes")
	sector := NewSector(area)
	if len(sector.animals) != 0 {
		t.Error("Animals count should be 0")
	}

	sector.GenerateAndAddAnimal()
	if len(sector.animals) != 1 {
		t.Error("Animals count should be 0")
	}
}

func TestFindAnimalByName(t *testing.T) {
	const animalType string = "fishes"
	area := NewArea(animalType)
	sector := NewSector(area)

	createAndAddAnimal := func(id int, name string) {
		an1 := &Animal{
			id:         id,
			animalType: animalType,
			name:       name,
		}
		sector.animals = append(sector.animals, an1)

	}

	createAndAddAnimal(0, "Jack")
	createAndAddAnimal(1, "Bob")
	createAndAddAnimal(2, "Carmen")
	createAndAddAnimal(3, "Emma")

	animal := sector.FindAnimalByName("Carmen")
	if animal == nil {
		t.Error("Carmen should be found")
	}

	animal = sector.FindAnimalByName("Louise")
	if animal != nil {
		t.Error("Louise shouldn't be found")
	}
}

func TestAnimalFeeding(t *testing.T) {
	area := NewArea("fishes")
	sector := NewSector(area)
	sector.GenerateAndAddAnimal()

	err := sector.utilitySpace.Feeding(0)
	if err != nil {
		t.Error("This animal should exist!")
	}

	err = sector.utilitySpace.Feeding(-1)
	if err == nil {
		t.Error("This animal doesn't exist!")
	}

	err = sector.utilitySpace.Feeding(1)
	if err == nil {
		t.Error("This animal doesn't exist!")
	}
}
