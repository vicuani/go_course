package filter

import (
	"log/slog"
)

type Aquarium interface {
	Size() int
	Animal() string
	SaltLevel() float64
	Contaminants() float64
	FilterSpeed() float64
	Cleaners() float64

	IncreaseFiltration(coef float64)
	DecreaseFiltration(coef float64)
	AddSalt(coef float64)
	AddCleaners(coef float64)
	CalculatePollution() float64
}

type Filter interface {
	Aquarium() Aquarium

	Adjust(pollutionLevel float64)
	AddSalt()
	AddCleaners()
}

type BasicFilter struct {
	aq Aquarium
}

func NewBasic(aq Aquarium) *BasicFilter {
	return &BasicFilter{
		aq: aq,
	}
}

func (f *BasicFilter) Aquarium() Aquarium {
	return f.aq
}

func (f *BasicFilter) Adjust(pollutionLevel float64) {
	// Just some formula
	adjustment := pollutionLevel * f.aq.FilterSpeed() / 10.0
	f.aq.IncreaseFiltration(adjustment)

	slog.Info("Adjusted filtration",
		"aquarium", f.aq.Animal(),
		"pollution level", pollutionLevel,
		"adjustment", adjustment,
	)
}

func (f *BasicFilter) AddSalt() {
	// Just some formula
	saltAdjustment := (0.5 - f.aq.SaltLevel()) * float64(f.aq.Size()) * 0.1
	if saltAdjustment > 0 {
		f.aq.AddSalt(saltAdjustment)
		slog.Info("Added salt",
			"aquarium", f.aq.Animal(),
			"salt level", f.aq.SaltLevel(),
			"added salt", saltAdjustment,
		)
	}
}

func (f *BasicFilter) AddCleaners() {
	// Just some formula
	cleanersAdjustment := (f.aq.Contaminants() - 0.3) * float64(f.aq.Size()) * 0.2
	if cleanersAdjustment > 0 {
		f.aq.AddCleaners(cleanersAdjustment)
		slog.Info("Added cleaners",
			"aquarium", f.aq.Animal(),
			"contaminants", f.aq.Contaminants(),
			"added cleaners", cleanersAdjustment,
		)
	}
}
