package animal

import (
	"testing"
)

func TestAnimal_InZone(t *testing.T) {
	a := NewAnimal(Deer, 120)

	t.Run("inZone", func(t *testing.T) {
		a.SetInZone(true)

		if !a.InZone() {
			t.Errorf("Expected animal to be in feeding zone, but it was not.")
		}
	})

	t.Run("notInZone", func(t *testing.T) {
		a.SetInZone(false)

		if a.InZone() {
			t.Errorf("Expected animal to be out of feeding zone, but it was in.")
		}
	})
}

func TestAnimalDetector_Detect(t *testing.T) {
	zone := Zone{
		Animals: []AnimalInterface{
			NewAnimal(Bear, 200),
			NewAnimal(Deer, 120),
		},
	}
	detector := &Detector{}

	zone.Animals[0].SetInZone(true)
	detectedAnimals := detector.Detect(&zone)

	if len(detectedAnimals) != 1 {
		t.Errorf("Expected 1 animal to be detected, but got %v", len(detectedAnimals))
	}
}
