package animal

import "fmt"

const FeederCapacity = 100.0

type Feeder struct {
	ID     int
	busy   bool
	volume float64
}

func NewFeeder(id int) *Feeder {
	return &Feeder{
		ID:     id,
		busy:   false,
		volume: FeederCapacity,
	}
}

func GenerateFeeders(n int) []*Feeder {
	var feeders []*Feeder
	for i := 0; i < n; i++ {
		feeder := NewFeeder(i)
		feeders = append(feeders, feeder)
	}
	return feeders
}

func (f *Feeder) free() {
	f.busy = false
}

func (f *Feeder) Refill() {
	f.volume = FeederCapacity
}

func (f *Feeder) IsBusy() bool {
	return f.busy
}

func (f *Feeder) IsEmpty() bool {
	return f.volume == 0.0
}

func (f *Feeder) CanBeUsed() bool {
	return !f.IsBusy() && !f.IsEmpty()
}

func (f *Feeder) Feed(an *Animal) error {
	defer f.free()

	if !f.CanBeUsed() {
		return fmt.Errorf("animal #%v cannot eat from feeder #%v", an.ID, f.ID)
	}

	f.busy = true
	return an.eat(f)
}
