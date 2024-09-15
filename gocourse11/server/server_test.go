package server

import (
	"testing"

	"github.com/vicuani/go_course/gocourse11/aquarium"
	"github.com/vicuani/go_course/gocourse11/filter"
)

func TestMonitorAndFilter(t *testing.T) {
	aq := aquarium.NewAquariumBuilder().
		SetSize(100).
		SetContaminants(0.4).
		Build()

	basicFilter := filter.NewBasicFilter(aq)
	srv := Server{}
	srv.AddAquarium(aq)
	srv.AddFilterSystem(basicFilter)

	srv.MonitorAndFilter()

	pollutionLevel := srv.CalculatePollution(aq)
	expectedPollution := 0.004
	if pollutionLevel != expectedPollution {
		t.Errorf("expected pollution level %f, got %f", expectedPollution, pollutionLevel)
	}
}
