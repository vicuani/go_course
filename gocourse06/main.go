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
	"math/rand/v2"
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse06/animal"
)

const iterationsCount = 10

const animalCount = 10
const feedersCount = 3
const maxEnclosuresAccess = 5

type EnclosureRequest struct {
	AnimalID     int
	IsOpenAction bool
	RespCh       chan bool
}

func enclosureController(reqCh <-chan EnclosureRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	accessCount := 0
	for req := range reqCh {
		if req.IsOpenAction {
			if accessCount < maxEnclosuresAccess {
				accessCount++
				req.RespCh <- true
				fmt.Printf("Access granted to enclosure for animal #%d (current accesses: %d)\n", req.AnimalID, accessCount)
			} else {
				req.RespCh <- false
				fmt.Printf("Access denied to enclosure for animal #%d, limit reached\n", req.AnimalID)
			}
		} else {
			accessCount--
			fmt.Printf("Enclosure closed by animal #%d (current accesses: %d)\n", req.AnimalID, accessCount)
		}
	}
}

func emulateAnimalChanges(an *animal.Animal, anCh chan<- *animal.Animal, logCh chan<- string, enclosureReqCh chan<- EnclosureRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	for range iterationsCount {
		an.RandomlyChangeIndicators()

		if an.IsHungry() {
			respCh := make(chan bool)
			enclosureReqCh <- EnclosureRequest{
				AnimalID:     an.ID,
				IsOpenAction: true,
				RespCh:       respCh,
			}
			accessGranted := <-respCh

			if accessGranted {
				anCh <- an
				logCh <- fmt.Sprintf("Animal #%v is hungry and got access to the enclosure, satiety = %v\n", an.ID, an.Satiety)

				enclosureReqCh <- EnclosureRequest{
					AnimalID:     an.ID,
					IsOpenAction: false,
				}
			} else {
				logCh <- fmt.Sprintf("Animal #%v is hungry but access to the enclosure was denied, satiety = %v\n", an.ID, an.Satiety)
			}
		}

		if an.HasCriticalValues() {
			logCh <- fmt.Sprintf("Animal #%v has critical value(s): health = %v, mood = %v, satiety = %v\n", an.ID, an.Health, an.Mood, an.Satiety)
		}

		time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.IntN(100)+50)))
	}
}

func handleFeeders(fCh chan *animal.Feeder, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case f, ok := <-fCh:
			if !ok {
				return
			}

			if f.IsEmpty() {
				f.Refill()
				fmt.Printf("Check-up for feeder #%v, it is empty, refill it\n", f.ID)
			}

			fCh <- f
		case <-doneCh:
			fmt.Println("Received done signal, exiting handleFeeders...")
			return
		}
		time.Sleep(time.Duration(time.Millisecond * time.Duration(10)))
	}
}

func handleHunger(fCh chan *animal.Feeder, anCh <-chan *animal.Animal, logCh chan<- string, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case an, ok := <-anCh:
			if !ok {
				return
			}

			fmt.Printf("Here are only hungry animals: #%v, satiety = %v\n", an.ID, an.Satiety)
			if an.IsHungry() {
				//	search for not empty feeder
				for {
					feeder, ok := <-fCh
					if !ok {
						return
					}
					feedRes := feeder.Feed(an)
					if feedRes != nil {
						logCh <- feedRes.Error()
					} else {
						fCh <- feeder
						break
					}
					fCh <- feeder
				}
			}
		case <-doneCh:
			fmt.Println("Received done signal, exiting handleHunger...")
			return
		}
	}
}

func mainMonitor(logCh <-chan string, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case log, ok := <-logCh:
			if !ok {
				return
			}
			fmt.Printf("Log: %v\n", log)
		case <-doneCh:
			fmt.Println("Received done signal, exiting mainMonitor...")
			return
		}
	}
}

func main() {
	animals := animal.GenerateAnimals(animalCount)
	feeders := animal.GenerateFeeders(feedersCount)

	fmt.Println("\nAnimals")
	for _, an := range animals {
		fmt.Println(*an)
	}

	fmt.Println("\nFeeders")
	for _, f := range feeders {
		fmt.Println(*f)
	}

	var wg sync.WaitGroup
	var anWg sync.WaitGroup
	doneCh := make(chan struct{})
	anCh := make(chan *animal.Animal, animalCount)
	enclosureReqCh := make(chan EnclosureRequest)

	fCh := make(chan *animal.Feeder, len(feeders))
	for _, f := range feeders {
		fCh <- f
	}

	wg.Add(1)
	go enclosureController(enclosureReqCh, &wg)

	wg.Add(1)
	go handleFeeders(fCh, doneCh, &wg)

	logCh := make(chan string)
	wg.Add(1)
	go mainMonitor(logCh, doneCh, &wg)

	fmt.Println("Let's go...")

	for _, an := range animals {
		anWg.Add(1)
		go func(an *animal.Animal) {
			defer anWg.Done()
			emulateAnimalChanges(an, anCh, logCh, enclosureReqCh, &wg)
		}(an)
	}

	wg.Add(1)
	go handleHunger(fCh, anCh, logCh, doneCh, &wg)

	go func() {
		anWg.Wait()
		close(anCh)
	}()

	wg.Wait()
	close(doneCh)
}
