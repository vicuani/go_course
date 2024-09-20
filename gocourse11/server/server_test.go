package server

import (
	"testing"

	"github.com/vicuani/go_course/gocourse11/aquarium"
	"github.com/vicuani/go_course/gocourse11/filter"
)

func TestMonitorAndFilter(t *testing.T) {
	aq := aquarium.NewBuilder().
		SetSize(100).
		SetContaminants(0.4).
		Build()

	basicFilter := filter.NewBasic(aq)
	srv := Server{}
	srv.AddFilter(basicFilter)

	srv.MonitorAndFilter()

	pollutionLevel := aq.CalculatePollution()
	expectedPollution := 0.004
	if pollutionLevel != expectedPollution {
		t.Errorf("expected pollution level %f, got %f", expectedPollution, pollutionLevel)
	}
}
