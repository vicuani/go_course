package server

import (
	"log/slog"

	"github.com/vicuani/go_course/gocourse11/filter"
)

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
