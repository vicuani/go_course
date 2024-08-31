package main

import "testing"

func TestNewSector(t *testing.T) {
	area := NewArea("fishes")
	sector := NewSector(area)

	if sector == nil {
		t.Error("Sector wasn't created")
		return
	}

	if sector.area != area {
		t.Errorf("Incorrect area: expected: %v, current: %v", area, sector.area)
	}

	if sector.utilitySpace == nil {
		t.Error("Utility space inside sector wasn't created")
	}
}

func TestCreateAndAddAnimal(t *testing.T) {
	animalType := "fishes"
	area := NewArea(animalType)
	sector := NewSector(area)

	id := 17
	name := "Nemo"
	sector.CreateAndAddAnimal(id, name)

	if len(sector.animals) != 1 {
		t.Error("Animal wasn't created and/or added")
	}

	addedAnimal := sector.animals[0]
	if addedAnimal == nil {
		t.Error("Animal wasn't created")
		return
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

func preconditions() *Sector {
	const animalType string = "mammals"
	area := NewArea(animalType)
	sector := NewSector(area)

	sector.CreateAndAddAnimal(0, "Jack")
	sector.CreateAndAddAnimal(1, "Bob")
	sector.CreateAndAddAnimal(2, "Carmen")
	sector.CreateAndAddAnimal(3, "Emma")

	return sector
}

func TestFindAnimalByName(t *testing.T) {
	sector := preconditions()

	animal := sector.FindAnimalByName("Carmen")
	if animal == nil {
		t.Error("Carmen should be found")
	}

	animal = sector.FindAnimalByName("Louise")
	if animal != nil {
		t.Error("Louise shouldn't be found")
	}
}

func TestGetAnimalIndex(t *testing.T) {
	sector := preconditions()

	animal := sector.FindAnimalByName("Bob")
	index := sector.GetAnimalIndex(animal)

	if index == - 1 {
		t.Error("This animal shouldn't have index -1!")
	}
}
