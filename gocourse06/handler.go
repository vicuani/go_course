package main

import (
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse06/animal"
)

func monitorSystem(logChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case logData, ok := <-logChan:
			if !ok {
				logger.Info("End monitoring...")
				return
			}

			logger.Info("[LOG]", "message", logData)
		default:
			time.Sleep(time.Duration(10 * time.Millisecond))
		}
	}
}

func handleFeeder(feederChan <-chan *animal.Feeder, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case f, ok := <-feederChan:
			if !ok {
				logger.Info("End handling feeders state...")
				return
			}

			if !f.IsEmpty() {
				logger.Info("Only empty feeders should be handled")
				break
			}

			logger.Info("Refilling feeder", "id", f.ID)
			f.Refill()

		default:
			time.Sleep(time.Duration(10 * time.Millisecond))
		}
	}
}

func handleHunger(feeders []*animal.Feeder, animalChan <-chan *animal.Animal, feederChan chan<- *animal.Feeder, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case an, ok := <-animalChan:
			if !ok {
				logger.Info("End handling hunger...")
				close(feederChan)
				return
			}

			if !an.IsHungry() {
				logger.Info("Only hungry animals should be handled by hunger", "id", an.ID)
				break
			}

			logger.Info("Handling hunger", "animal", an)
			for _, feeder := range feeders {
				res := feeder.Feed(an)
				if feeder.IsEmpty() {
					feederChan <- feeder
				}

				if res {
					break
				}
			}
		default:
			time.Sleep(time.Duration(10 * time.Millisecond))
		}
	}
}
