package main

import (
	"testing"
)

func TestFeederObservers(t *testing.T) {
	lChan := make(chan string, 10)
	fdr := feeder{
		ordinaryFood: generateOrdinaryFood(),
	}

	animals := []*animal{
		newAnimal(1, lChan),
		newAnimal(2, lChan),
	}

	for _, an := range animals {
		fdr.addEater(an)
	}

	strategy := &ordinaryFeedingStrategy{fdr: &fdr}
	fdr.setStrategy(strategy)

	fdr.feedAll()

	for i := range len(animals) {
		select {
		case msg := <-lChan:
			if msg == "" {
				t.Error("Expected non-empty message")
			}
		default:
			t.Errorf("Expected message from animal %v, but none received", i+1)
		}
	}
}
