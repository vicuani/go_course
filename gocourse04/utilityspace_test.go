package main

import "testing"

func TestNewUtilitySpace(t *testing.T) {
	stubSector := &Sector{
		area: NewArea("fishes"),
	}
	utilitySpace := NewUtilitySpace(stubSector)

	if utilitySpace == nil {
		t.Fatal("Utility space wasn't created")
	}

	if utilitySpace.sector != stubSector {
		t.Errorf("Incorrect sector: expected: %v, current: %v", stubSector, utilitySpace.sector)
	}
}

func TestAnimalFeeding(t *testing.T) {
	area := NewArea("fishes")
	sector := NewSector(area)
	sector.AddAnimal(sector.NewRandomAnimal())
	animal1 := sector.animals[0]

	err := sector.utilitySpace.Feed(animal1)
	if err != nil {
		t.Error("This animal should exist!")
	}

	animal2 := NewAnimal(33, "fishes", "Eric")
	err = sector.utilitySpace.Feed(animal2)
	if err == nil {
		t.Error("This animal doesn't exist!")
	}
}
