package server

import (
	"log/slog"

	"github.com/vicuani/go_course/gocourse11/filter"
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

type Server struct {
	filters   []filter.Filter
}

func (s *Server) AddFilter(f filter.Filter) {
	s.filters = append(s.filters, f)
}

func (s *Server) MonitorAndFilter() {
	for _, f := range s.filters {
		pollutionLevel := f.Aquarium().CalculatePollution()

		slog.Info("Monitoring aquarium",
			"animal", f.Aquarium().Animal(),
			"pollutionLevel", pollutionLevel,
		)

		f.Adjust(pollutionLevel)
		f.AddSalt()
		f.AddCleaners()
	}
}
