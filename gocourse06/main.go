/*

«Загальна система розумного зоопарку»

Концепція

Створити програму для управління розумним зоопарком,
де декілька горутин виконують різні завдання, такі як
	моніторинг стану тварин,
	керування доступом до вольєрів та
	управління кормушками.
Програма має активно використовувати канали для комунікації між горутинами, уникаючи «race conditions» і «deadlocks».

Завдання

Моніторинг стану тварин:
Створіть горутину для кожної тварини в зоопарку.
Кожна горутина збирає дані про стан тварини
(наприклад, рівень здоров'я, голод, настрій)
і відправляє їх через канал до центральної системи моніторингу.

Керування доступом до вольєрів:
Імплементуйте горутину, яка
	контролює доступ до вольєрів, використовуючи канали для отримання запитів на відкриття/закриття.


Управління кормушками:
Розробіть горутини для управління автоматичними кормушками, які
	відправляють статус кормушки (порожня/повна) через канал.

Умови виконання

Уникнення «Race Conditions»:
забезпечте, щоб спільні ресурси (наприклад, дані про стан тварин) були захищені від одночасного доступу декількома горутинами.
Використовуйте канали для синхронізації доступу.

Управління «Deadlocks»:
уважно використовуйте блокування та канали,
щоб уникнути взаємних блокувань між горутинами.

Логування та моніторинг:
реалізуйте систему логування, яка
фіксує важливі події у системі, наприклад, коли
	тварина потребує уваги або
	коли кормушка порожня.

Тестування: напишіть модульні тести для перевірки коректності взаємодії між горутинами та уникнення «race conditions» та «deadlocks».

*/

package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse06/animal"
)

const (
	iterationsCount     = 10
	maxEnclosuresAccess = 5
	animalCount         = 20
	feedersCount        = 3
)

var logger *slog.Logger

type EnclosureRequest struct {
	animalID     int
	isOpenAction bool
	respChan     chan bool
}

func enclosureController(enclosureReqChan <-chan EnclosureRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	accessCount := 0
	for req := range enclosureReqChan {
		if req.isOpenAction {
			if accessCount < maxEnclosuresAccess {
				accessCount++
				req.respChan <- true
				logger.Info("Access granted to enclosure for animal", "id", req.animalID, "current accesses", accessCount)
			} else {
				req.respChan <- false
				logger.Info("Access denied to enclosure (limit reached) for animal", "id", req.animalID)
			}
		} else {
			accessCount--
			logger.Info("Enclosure closed by animal", "id", req.animalID, "current accesses", accessCount)
		}
	}
}

func emulateAnimalChanges(an *animal.Animal, animalChan chan<- *animal.Animal, enclosureReqCh chan<- EnclosureRequest, logChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range iterationsCount {
		logger.Info("*** Iteration", "#", i, "for animal", an.ID)
		an.RandomlyChangeIndicators()
		if an.IsHungry() {

			respCh := make(chan bool)
			enclosureReqCh <- EnclosureRequest{
				animalID:     an.ID,
				isOpenAction: true,
				respChan:     respCh,
			}
			accessGranted := <-respCh
			close(respCh)

			if accessGranted {
				enclosureReqCh <- EnclosureRequest{
					animalID:     an.ID,
					isOpenAction: false,
				}

				animalChan <- an
			} else {
				logChan <- fmt.Sprintf("Animal #%v is hungry but access to the enclosure was denied, satiety = %v\n", an.ID, an.Satiety())
			}

		}
		if an.HasCriticalValues() {
			logChan <- string("animal has critical values:" + an.String())
		}
		time.Sleep(time.Duration(time.Millisecond * time.Duration(100)))
	}
}

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

	enclosureReqChan := make(chan EnclosureRequest)
	animalChan := make(chan *animal.Animal, animalCount)
	feederChan := make(chan *animal.Feeder)
	logChan := make(chan string)

	var wg sync.WaitGroup

	var animals []*animal.Animal
	for i := 0; i < animalCount; i++ {
		animals = append(animals, animal.NewAnimal(i, logger))
	}

	var feeders []*animal.Feeder
	for i := 0; i < feedersCount; i++ {
		feeder := animal.NewFeeder(i, logger)
		feeders = append(feeders, feeder)
	}

	wg.Add(1)
	go monitorSystem(logChan, &wg)

	wg.Add(1)
	go handleFeeder(feederChan, &wg)

	wg.Add(1)
	go handleHunger(feeders, animalChan, feederChan, &wg)

	wg.Add(1)
	go enclosureController(enclosureReqChan, &wg)

	{
		var animalWg sync.WaitGroup
		for _, an := range animals {
			animalWg.Add(1)
			go emulateAnimalChanges(an, animalChan, enclosureReqChan, logChan, &animalWg)
		}

		animalWg.Wait()
		close(animalChan)
		close(enclosureReqChan)
		close(logChan)
	}

	wg.Wait()
	logger.Info("End")
}
