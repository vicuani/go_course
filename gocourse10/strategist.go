package main

import (
	"math/rand/v2"
	"time"
)

type ordinaryFeedingStrategy struct {
	fdr *feeder
}

func (ofs *ordinaryFeedingStrategy) getFood() []food {
	randOrdFood := ofs.fdr.ordinaryFood[rand.IntN(len(ofs.fdr.ordinaryFood))]
	return []food{randOrdFood}
}

type weekendFeedingStrategy struct {
	fdr *feeder
}

func (wfs *weekendFeedingStrategy) getFood() []food {
	randOrdFood := wfs.fdr.ordinaryFood[rand.IntN(len(wfs.fdr.ordinaryFood))]
	randWkndFood := wfs.fdr.weekendFood[rand.IntN(len(wfs.fdr.weekendFood))]
	return []food{randOrdFood, randWkndFood}
}

type strategist struct {
	ofs *ordinaryFeedingStrategy
	wfs *weekendFeedingStrategy
}

func newStrategist(fdr *feeder) *strategist {
	return &strategist{
		ofs: &ordinaryFeedingStrategy{fdr: fdr},
		wfs: &weekendFeedingStrategy{fdr: fdr},
	}
}

func (s *strategist) getCorrectStrategy(day time.Weekday) feedingStrategy {
	if day == time.Saturday || day == time.Sunday {
		return s.wfs
	}
	return s.ofs
}
