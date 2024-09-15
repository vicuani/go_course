package server

import (
	"log/slog"

	"github.com/vicuani/go_course/gocourse11/aquarium"
	"github.com/vicuani/go_course/gocourse11/filter"
)

type Server struct {
	Aquariums []aquarium.AquariumInterface
	Filters   []filter.FilterSystem
}

func (s *Server) AddAquarium(a aquarium.AquariumInterface) {
	s.Aquariums = append(s.Aquariums, a)
}

func (s *Server) AddFilterSystem(f filter.FilterSystem) {
	s.Filters = append(s.Filters, f)
}

func (s *Server) MonitorAndFilter() {
	for _, f := range s.Filters {
		pollutionLevel := s.CalculatePollution(f.Aquarium())

		slog.Info("Monitoring aquarium",
			"animal", f.Aquarium().Animal(),
			"pollutionLevel", pollutionLevel,
		)

		f.Adjust(pollutionLevel)
		f.AddSalt()
		f.AddCleaners()
	}
}

func (s *Server) CalculatePollution(a aquarium.AquariumInterface) float64 {
	return a.Contaminants() / float64(a.Size())
}
