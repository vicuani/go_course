package feeder

import (
	"log/slog"
	"os"
	"testing"

	"github.com/vicuani/go_course/gocourse09/animal"
)

func TestFeeder_Refill(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	f := NewFeeder(0, logger)
	f.Refill(50)

	if got := f.Stock(); got != 50 {
		t.Errorf("Expected stock to be 50, but got %v", got)
	}

	f.Refill(30)

	if got := f.Stock(); got != 80 {
		t.Errorf("Expected stock to be 80, but got %v", got)
	}
}

func TestFeeder_CalculateFood(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	f := NewFeeder(0, logger)

	deer := animal.NewAnimal(animal.Deer, 120)
	bracket := f.calculateFood(deer)

	if bracket.amount != 4*deer.Weight()/100 {
		t.Errorf("Expected amount to be %v, but got %v", 4*deer.Weight()/100, bracket.amount)
	}
}
