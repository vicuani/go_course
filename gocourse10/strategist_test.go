package main

import (
	"testing"
	"time"
)

func TestGetCorrectStrategy(t *testing.T) {
	fdr := feeder{
		ordinaryFood: generateOrdinaryFood(),
		weekendFood:  generateWeekendFood(),
	}

	strategist := newStrategist(&fdr)

	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday}
	weekends := []time.Weekday{time.Saturday, time.Sunday}

	t.Run("weekdays", func(t *testing.T) {
		for _, day := range weekdays {
			strategy := strategist.getCorrectStrategy(day)
			food := strategy.getFood()

			if len(food) != 1 {
				t.Errorf("Expected ordinary food count of 1, got %v", len(food))
			}
		}
	})

	t.Run("weekends", func(t *testing.T) {
		for _, day := range weekends {
			strategy := strategist.getCorrectStrategy(day)
			food := strategy.getFood()

			if len(food) != 2 {
				t.Errorf("Expected weekend food count of 2, got %v", len(food))
			}
		}
	})
}
