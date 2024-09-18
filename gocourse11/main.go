package main

import (
	"github.com/vicuani/go_course/gocourse11/aquarium"
	"github.com/vicuani/go_course/gocourse11/filter"
	"github.com/vicuani/go_course/gocourse11/server"
)

func main() {
	aquariums := []*aquarium.Aquarium{
		aquarium.NewBuilder().
			SetSize(100).
			SetAnimal("Salmon").
			SetSaltLevel(0.3).
			SetContaminants(0.4).
			SetFilterSpeed(2.0).
			SetCleaners(0.5).
			Build(),

		aquarium.NewBuilder().
			SetSize(200).
			SetAnimal("Shrimp").
			SetSaltLevel(0.2).
			SetContaminants(0.1).
			SetFilterSpeed(1.5).
			SetCleaners(0.7).
			Build(),

		aquarium.NewBuilder().
			SetSize(2000).
			SetAnimal("Crocodile").
			SetSaltLevel(0.4).
			SetContaminants(0.2).
			SetFilterSpeed(5.0).
			SetCleaners(2.0).
			Build(),
	}

	centralServer := server.Server{}

	for _, aq := range aquariums {
		centralServer.AddFilter(filter.NewBasic(aq))
	}

	centralServer.MonitorAndFilter()
}
