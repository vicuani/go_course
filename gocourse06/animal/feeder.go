package animal

import (
	"fmt"
	"log/slog"
	"sync"
)

const FeederCapacity = 100

type Feeder struct {
	ID int

	parametersMu sync.Mutex
	volume       int

	logger *slog.Logger
}

func NewFeeder(id int, logger *slog.Logger) *Feeder {
	return &Feeder{
		ID:     id,
		volume: FeederCapacity,
		logger: logger,
	}
}

func (f *Feeder) String() string {
	return fmt.Sprintf("id = %v, volume = %v", f.ID, f.Volume())
}

func (f *Feeder) Volume() int {
	defer f.parametersMu.Unlock()

	f.parametersMu.Lock()
	return f.volume
}

func (f *Feeder) SetVolume(v int) {
	f.parametersMu.Lock()
	f.volume = v
	f.parametersMu.Unlock()
}

func (f *Feeder) IsEmpty() bool {
	return f.Volume() == 0
}

func (f *Feeder) Refill() {
	f.SetVolume(FeederCapacity)
}

func (f *Feeder) Feed(an *Animal) bool {
	if f.IsEmpty() {
		f.logger.Info("Feeder cannot feed animal, it's empty", "id", f.ID)
		return false
	}

	eaten := min(MaxIndicatorValue-an.Satiety(), f.Volume())
	an.eat(eaten)
	f.SetVolume(f.Volume() - eaten)
	f.logger.Info("Feeder", "id", f.ID, "feed animal", an.ID, "left volume", f.Volume())
	return true
}
