package filter

import (
	"math"
	"testing"

	"github.com/vicuani/go_course/gocourse11/aquarium"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestAdjustFiltration(t *testing.T) {
	aq := aquarium.NewBuilder().
		SetSize(100).
		SetAnimal("Fish").
		SetFilterSpeed(1.5).
		Build()

	basicFilter := NewBasic(aq)
	pollutionLevel := 0.6

	basicFilter.Adjust(pollutionLevel)
	expectedFiltration := 1.59
	if !almostEqual(aq.FilterSpeed(), expectedFiltration) {
		t.Errorf("expected filtration adjustment %f, got %f", expectedFiltration, aq.FilterSpeed())
	}
}

func TestAddSalt(t *testing.T) {
	aq := aquarium.NewBuilder().
		SetSize(100).
		SetSaltLevel(0.2).
		SetContaminants(0.4).
		Build()

	basicFilter := NewBasic(aq)
	basicFilter.AddSalt()

	expectedSalt := 3.2
	if !almostEqual(aq.SaltLevel(), expectedSalt) {
		t.Errorf("expected salt addition %f, got %f", expectedSalt, aq.SaltLevel())
	}
}

func TestAddCleaners(t *testing.T) {
	aq := aquarium.NewBuilder().
		SetSize(200).
		SetSaltLevel(0.1).
		SetContaminants(0.5).
		Build()

	basicFilter := NewBasic(aq)
	basicFilter.AddCleaners()

	expectedCleaners := 8.0
	if !almostEqual(aq.Cleaners(), expectedCleaners) {
		t.Errorf("expected cleaners addition %f, got %f", expectedCleaners, aq.Cleaners())
	}
}
