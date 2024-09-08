package feeder

import (
	"testing"

	"github.com/vicuani/go_course/gocourse09/animal"
)

func TestFeeder_Refill(t *testing.T) {
	f := NewFeeder(0)
	f.Refill(50)

	if f.Stock() != 50 {
		t.Errorf("Expected stock to be 50, but got %v", f.Stock())
	}

	f.Refill(30)

	if f.Stock() != 80 {
		t.Errorf("Expected stock to be 80, but got %v", f.Stock())
	}
}

func TestFeeder_CalculateFood(t *testing.T) {
	f := NewFeeder(0)

	deer := animal.NewAnimal(animal.Deer, 120)
	bracket := f.calculateFood(deer)

	if bracket.amount != 4*deer.Weight/100 {
		t.Errorf("Expected amount to be %v, but got %v", 4*deer.Weight/100, bracket.amount)
	}
}
