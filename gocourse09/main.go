package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse09/animal"
	"github.com/vicuani/go_course/gocourse09/feeder"
)

const locationIterations = 10

func simulateAnimalLocations(fChan chan<- bool, zone *animal.Zone, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < locationIterations; i++ {
		fmt.Printf("\nLocation iteration #%v\n", i)
		needToFeed := false
		for _, an := range zone.Animals {
			inZone := rand.IntN(2) == 1
			if inZone != an.InZone() {
				if inZone {
					fmt.Printf("%s appeared in zone\n", an.Type)
				} else {
					fmt.Printf("%s left the zone\n", an.Type)
				}
				an.SetInZone(inZone)
				needToFeed = true
			}
		}

		if needToFeed {
			fChan <- true
		}

		time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)+100))
	}
	close(fChan)
}

func handleRefiller(lsChan chan bool, f *feeder.Feeder, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case _, ok := <-lsChan:
			if !ok {
				fmt.Println("Low stock channel is closed, exiting from handleRefiller...")
				return
			}

			f.Refill(100)

		default:
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)+50))
		}
	}
}

func handleFeeder(lsChan chan bool, fChan <-chan bool, f *feeder.Feeder, zone *animal.Zone, wg *sync.WaitGroup) {
	defer wg.Done()

	d := &animal.AnimalDetector{}
	for {
		select {
		case _, ok := <-fChan:
			if !ok {
				fmt.Println("Feeder channel is closed, exiting from handleFeeder...")
				close(lsChan)
				return
			}

			ans := d.Detect(zone)
			f.Feed(lsChan, ans)

		default:
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)+50))
		}
	}
}

func main() {
	zone := animal.GenerateZone()

	fChan := make(chan bool, locationIterations)
	lsChan := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go simulateAnimalLocations(fChan, zone, &wg)

	f := feeder.NewFeeder(50)
	wg.Add(1)
	go handleRefiller(lsChan, f, &wg)

	wg.Add(1)
	go handleFeeder(lsChan, fChan, f, zone, &wg)

	wg.Wait()
}
