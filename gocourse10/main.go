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

type observer interface {
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

func generateAnimals(lChan chan string) []*animal {
	var animals []*animal
	for i := range animalCount {
		animals = append(animals, newAnimal(i, lChan))
	}

	return animals
}

func main() {
	fdr := feeder{
		ordinaryFood: generateOrdinaryFood(),
		weekendFood:  generateWeekendFood(),
	}

	lChan := make(chan string)
	animals := generateAnimals(lChan)
	for _, an := range animals {
		fdr.addObserver(an)
	}

	nextDay := func(day time.Time) time.Time {
		return day.AddDate(0, 0, 1)
	}

	d := time.Now()
	strategist := newStrategist(&fdr)
	fdr.setStrategy(strategist.getCorrectStrategy(d.Weekday()))

	var wg sync.WaitGroup
	logger := logger{}
	wg.Add(1)
	go logger.printLogs(lChan, &wg)

	for range daysCount {
		fmt.Printf("\nCurrent weekday is %v\n", d.Weekday())

		strategy := strategist.getCorrectStrategy(d.Weekday())
		if strategy != fdr.strategy() {
			fdr.setStrategy(strategy)
		}

		fdr.feedAll()
		d = nextDay(d)
	}
	close(lChan)

	wg.Wait()
}
