package main

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse09/animal"
	"github.com/vicuani/go_course/gocourse09/feeder"
)

const locationIterations = 10

var logger *slog.Logger

func simulateAnimalLocations(feederChan chan<- bool, zone *animal.Zone, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range locationIterations {
		logger.Info(fmt.Sprintf("Location iteration #%v", i))
		needToFeed := false
		for _, an := range zone.Animals {
			inZone := rand.IntN(2) == 1
			if inZone != an.InZone() {
				if inZone {
					logger.Info(fmt.Sprintf("%s appeared in zone", an.Type()))
				} else {
					logger.Info(fmt.Sprintf("%s left the zone", an.Type()))
				}
				an.SetInZone(inZone)
				needToFeed = true
			}
		}

		if needToFeed {
			feederChan <- true
		}

		time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)+100))
	}
	close(feederChan)
}

func handleRefiller(lowStockChan chan bool, f feeder.FeederInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case _, ok := <-lowStockChan:
			if !ok {
				logger.Info("Low stock channel is closed, exiting from handleRefiller...")
				return
			}

			f.Refill(100)

		default:
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)+50))
		}
	}
}

func handleFeeder(lowStockChan chan bool, feederChan <-chan bool, f feeder.FeederInterface, zone *animal.Zone, wg *sync.WaitGroup) {
	defer wg.Done()

	d := &animal.Detector{}
	for {
		select {
		case _, ok := <-feederChan:
			if !ok {
				logger.Info("Feeder channel is closed, exiting from handleFeeder...")
				close(lowStockChan)
				return
			}

			ans := d.Detect(zone)
			f.Feed(lowStockChan, ans)

		default:
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)+50))
		}
	}
}

func GenerateAnimals() []animal.AnimalInterface {
	return []animal.AnimalInterface{
		animal.NewAnimal(animal.Bear, 200),
		animal.NewAnimal(animal.Deer, 120),
		animal.NewAnimal(animal.Lion, 150),
		animal.NewAnimal(animal.Wolf, 50),
	}
}

func GenerateZone() *animal.Zone {
	return &animal.Zone{Animals: GenerateAnimals()}
}

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

	zone := GenerateZone()

	feederChan := make(chan bool, locationIterations)
	lowStockChan := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go simulateAnimalLocations(feederChan, zone, &wg)

	f := feeder.NewFeeder(50, logger)
	wg.Add(1)
	go handleRefiller(lowStockChan, f, &wg)

	wg.Add(1)
	go handleFeeder(lowStockChan, feederChan, f, zone, &wg)

	wg.Wait()
}
