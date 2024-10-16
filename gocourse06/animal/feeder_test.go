package animal

import (
	"io"
	"log/slog"
	"testing"
)

func TestFeed(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	feeder := NewFeeder(1, logger)
	if feeder.Volume() != 100 {
		t.Errorf("Feeder volume expected to be 100, got %d", feeder.Volume())
	}

	animal := NewAnimal(1, logger)
	animal.SetSatiety(80)
	feedSuccess := feeder.Feed(animal)
	if !feedSuccess {
		t.Errorf("Expected feeder to successfully feed animal")
	}

	if feeder.Volume() != 80 {
		t.Errorf("Expected feeder volume to decrease after feeding")
	}

	animal.SetSatiety(80)
	feeder.SetVolume(0)
	feedSuccess = feeder.Feed(animal)
	if feedSuccess {
		t.Errorf("Feeder should not feed when empty")
	}

	if feeder.Volume() != 0 {
		t.Errorf("Feeder volume should be zero")
	}
}
