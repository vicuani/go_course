package main

import "testing"

func TestNewSector(t *testing.T) {
	area := NewArea("reptiles")
	sector := NewSector(area)

	if sector == nil {
		t.Fatal("Sector wasn't created")
	}

	if sector.area != area {
		t.Errorf("Incorrect area: expected: %v, current: %v", area, sector.area)
	}

	if sector.utilitySpace == nil {
		t.Error("Utility space inside sector wasn't created")
	}
}

func TestNewAnimal(t *testing.T) {
	animalType := "fishes"
	area := NewArea(animalType)
	sector := NewSector(area)

	id := 17
	name := "Nemo"
	sector.AddAnimal(NewAnimal(id, animalType, name))

	if len(sector.animals) != 1 {
		t.Fatal("Animal wasn't created and/or added")
	}

	addedAnimal := sector.animals[0]
	if addedAnimal == nil {
		t.Fatal("Animal wasn't created")
	}

	if addedAnimal.id != id {
		t.Errorf("Incorrect id: expected: %v, current: %v", id, addedAnimal.id)
	}

	if addedAnimal.animalType != animalType {
		t.Errorf("Incorrect animal type: expected: %v, current: %v", animalType, addedAnimal.animalType)
	}

	if addedAnimal.name != name {
		t.Errorf("Incorrect name: expected: %v, current: %v", name, addedAnimal.name)
	}
}

func TestFindAnimalByName(t *testing.T) {
	area := NewArea("mammals")
	sector := NewSector(area)

	sector.AddAnimal(sector.NewAnimal("Jack"))
	sector.AddAnimal(sector.NewAnimal("Bob"))
	sector.AddAnimal(sector.NewAnimal("Carmen"))
	sector.AddAnimal(sector.NewAnimal("Emma"))

	t.Run("positive", func(t *testing.T) {
		animal := sector.FindAnimalByName("Carmen")
		if animal == nil {
			t.Error("Carmen should be found")
		}
	})

	t.Run("negative", func(t *testing.T) {
		animal := sector.FindAnimalByName("Louise")
		if animal != nil {
			t.Error("Louise shouldn't be found")
		}
	})
}

func TestAnimalIndex(t *testing.T) {
	area := NewArea("birds")
	sector := NewSector(area)

	sector.AddAnimal(sector.NewAnimal("Jack"))
	sector.AddAnimal(sector.NewAnimal("Bob"))
	sector.AddAnimal(sector.NewAnimal("Carmen"))
	sector.AddAnimal(sector.NewAnimal("Emma"))

	animal := sector.FindAnimalByName("Bob")
	index := sector.AnimalIndex(animal)

	if index == -1 {
		t.Error("This animal shouldn't have index -1!")
	}
}
