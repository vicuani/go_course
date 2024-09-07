package main

import "testing"

func TestCreateAnimal(t *testing.T) {
	id := 4
	name := "Carl"
	animalType := "dinosaur"
	an := NewAnimal(id, animalType, name)

	if an == nil {
		t.Fatal("Animal wasn't created")
	}

	if an.id != id {
		t.Errorf("Incorrect id: expected: %v, current: %v", id, an.id)
	}

	if an.animalType != animalType {
		t.Errorf("Incorrect animal type: expected: %v, current: %v", animalType, an.animalType)
	}

	if an.name != name {
		t.Errorf("Incorrect name: expected: %v, current: %v", name, an.name)
	}
}

func TestNewArea(t *testing.T) {
	animalType := "octopus"
	area := NewArea(animalType)

	if area == nil {
		t.Fatal("Area wasn't created")
	}

	if area.animalType != animalType {
		t.Errorf("Incorrect animal type: expected: %v, current: %v", animalType, area.animalType)
	}
}
