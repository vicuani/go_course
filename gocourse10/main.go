package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	animalCount = 5
	daysCount   = 8
)

type eater interface {
	eat([]food)
}

type food interface {
	name() string
	calories() int
}

type feedingStrategy interface {
	getFood() []food
}

func generateOrdinaryFood() []food {
	return []food{
		&sandwich{},
		&cheeseSandwich{},
		&potato{},
	}
}

func generateWeekendFood() []food {
	return []food{
		&cake{},
		&icecream{},
		&pizza{},
	}
}

func main() {
	fdr := feeder{
		ordinaryFood: generateOrdinaryFood(),
		weekendFood:  generateWeekendFood(),
	}

	lChan := make(chan string)
	for i := range animalCount {
		fdr.addEater(newAnimal(i, lChan))
	}

	nextDay := func(day time.Time) time.Time {
		return day.AddDate(0, 0, 1)
	}

	d := time.Now()
	last := d.AddDate(0, 0, daysCount)
	strategist := newStrategist(&fdr)
	fdr.setStrategy(strategist.getCorrectStrategy(d.Weekday()))

	var wg sync.WaitGroup
	logger := logger{}
	wg.Add(1)
	go logger.printLogs(lChan, &wg)

	for ; d.Before(last); d = nextDay(d) {
		fmt.Printf("\nCurrent weekday is %v\n", d.Weekday())

		strategy := strategist.getCorrectStrategy(d.Weekday())
		if strategy != fdr.strategy() {
			fdr.setStrategy(strategy)
		}

		fdr.feedAll()
	}
	close(lChan)

	wg.Wait()
}
