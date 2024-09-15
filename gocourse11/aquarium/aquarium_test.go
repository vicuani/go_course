package aquarium

import (
	"testing"
)

func TestAquariumBuilder(t *testing.T) {
	aq := NewAquariumBuilder().
		SetSize(100).
		SetAnimal("Salmon").
		SetSaltLevel(0.3).
		SetContaminants(0.4).
		SetFilterSpeed(2.0).
		SetCleaners(0.5).
		Build()

	if aq.Size() != 100 {
		t.Errorf("expected size 100, got %d", aq.Size())
	}
	if aq.Animal() != "Salmon" {
		t.Errorf("expected animal Salmon, got %s", aq.Animal())
	}
	if aq.SaltLevel() != 0.3 {
		t.Errorf("expected salt level 0.3, got %f", aq.SaltLevel())
	}
	if aq.Contaminants() != 0.4 {
		t.Errorf("expected contaminants 0.4, got %f", aq.Contaminants())
	}
	if aq.FilterSpeed() != 2.0 {
		t.Errorf("expected filter speed 2.0, got %f", aq.FilterSpeed())
	}
	if aq.Cleaners() != 0.5 {
		t.Errorf("expected cleaners 0.5, got %f", aq.Cleaners())
	}
}
