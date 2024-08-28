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

func monitorAnimalState(ch <-chan *animal.Animal, logChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case an, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Monitoring animal: %v\n", *an)
			if an.HasCriticalValues() {
				logChannel <- fmt.Sprintf("Animal %v needs attention! Health: %v, Hunger: %v, Mood: %v", an, an.Health, an.Hunger, an.Mood)
			}
		case <-time.After(time.Second * 5):
			return
		}
	}
}

func controlEnclosureAccess(ch <-chan *animal.Enclosure, logChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case en, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Controlling enclosure: %v: %v\n", en.ID, en.IsOpened)
			if !en.IsOpened {
				logChannel <- fmt.Sprintf("Enclosure is closed: %v", en.ID)
			}
		case <-time.After(time.Second * 5):
			return
		}
	}

}

func controlFeeder(ch <-chan *animal.Feeder, logChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case f, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Controlling feeder %v: %v\n", f.ID, f.IsEmpty)
			if f.IsEmpty {
				logChannel <- fmt.Sprintf("Feeder is empty: %v", f.ID)
			}
		case <-time.After(time.Second * 5):
			return
		}
	}
}

func emulateAnimalChanges(animals []*animal.Animal, animalChannel chan<- *animal.Animal) {
	const iterations = 10
	for i := 0; i < iterations; i++ {
		randAn := animals[rand.IntN(len(animals))]
		randAn.RandomlyChangeIndicators()
		animalChannel <- randAn
		time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.IntN(100))))
	}
	// close(animalChannel)
}

func emulateEnclosureChanges(enclosures []*animal.Enclosure, enclosureChannel chan<- *animal.Enclosure) {
	const iterations = 10
	for i := 0; i < iterations; i++ {
		randEn := enclosures[rand.IntN(len(enclosures))]
		randEn.IsOpened = rand.IntN(2) == 1
		enclosureChannel <- randEn
		time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.IntN(100))))
	}
	// close(enclosureChannel)
}

func emulateFeederChanges(feeders []*animal.Feeder, feederChannel chan<- *animal.Feeder) {
	const iterations = 10
	for i := 0; i < iterations; i++ {
		randF := feeders[rand.IntN(len(feeders))]
		randF.IsEmpty = rand.IntN(2) == 1
		feederChannel <- randF
		time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.IntN(100))))
	}
	// close(feederChannel)
}

func mainMonitor(animalChannel <-chan *animal.Animal, enclosureChannel <-chan *animal.Enclosure, feederChannel <-chan *animal.Feeder, logChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case state, ok := <-animalChannel:
			if !ok {
				animalChannel = nil
			} else {
				fmt.Printf("Animal state: Health=%v, Hunger=%v, Mood=%v\n", state.Health, state.Hunger, state.Mood)
			}
		case request, ok := <-enclosureChannel:
			if !ok {
				enclosureChannel = nil
			} else {
				fmt.Printf("Enclosure request: Enclosure %v %v\n", request.ID, request.IsOpened)
			}
		case status, ok := <-feederChannel:
			if !ok {
				feederChannel = nil
			} else {
				fmt.Printf("Feeder %v is %v\n", status.ID, status.IsEmpty)
			}
		case logEntry, ok := <-logChannel:
			if !ok {
				logChannel = nil
			} else {
				fmt.Println(logEntry)
			}
		}
		if animalChannel == nil && enclosureChannel == nil && feederChannel == nil && logChannel == nil {
			return
		}
	}
}

func main() {
	animals := animal.GenerateAnimals(10)
	enclosures := animal.GenerateEnclosures(5)
	feeders := animal.GenerateFeeders(3)

	fmt.Println("\nAnimals")
	for _, an := range animals {
		fmt.Println(*an)
	}

	fmt.Println("\nEnclosures")
	for _, en := range enclosures {
		fmt.Println(*en)
	}

	fmt.Println("\nFeeders")
	for _, f := range feeders {
		fmt.Println(*f)
	}

	fmt.Println("\nThe chaos starts right now...\n")

	animalChannel := make(chan *animal.Animal, len(animals))
	enclosureChannel := make(chan *animal.Enclosure, len(enclosures))
	feederChannel := make(chan *animal.Feeder, len(feeders))
	logChannel := make(chan string)

	var wg sync.WaitGroup

	go emulateAnimalChanges(animals, animalChannel)
	go emulateEnclosureChanges(enclosures, enclosureChannel)
	go emulateFeederChanges(feeders, feederChannel)

	for _, an := range animals {
		wg.Add(1)
		animalChannel <- an
		go monitorAnimalState(animalChannel, logChannel, &wg)
	}

	for _, en := range enclosures {
		wg.Add(1)
		enclosureChannel <- en
		go controlEnclosureAccess(enclosureChannel, logChannel, &wg)
	}

	for _, f := range feeders {
		wg.Add(1)
		feederChannel <- f
		go controlFeeder(feederChannel, logChannel, &wg)
	}

	wg.Add(1)
	go mainMonitor(animalChannel, enclosureChannel, feederChannel, logChannel, &wg)

	wg.Wait()
	fmt.Println("\nEnd")
}