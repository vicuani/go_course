package feeder

import (
	"fmt"
	"sync"

	"github.com/vicuani/go_course/gocourse09/animal"
)

type FoodBracket struct {
	animalType animal.AnimalType
	amount     int
}

type Feeder struct {
	stockMu sync.Mutex
	stock   int
}

func NewFeeder(initialStock int) *Feeder {
	return &Feeder{
		stock: initialStock,
	}
}

func (fd *Feeder) Stock() int {
	defer fd.stockMu.Unlock()

	fd.stockMu.Lock()
	return fd.stock
}

func (fd *Feeder) SetStock(v int) {
	fd.stockMu.Lock()
	fd.stock = v
	fd.stockMu.Unlock()
}

func (fd *Feeder) Feed(lsChan chan<- bool, animals []*animal.Animal) {
	for _, an := range animals {
		food := fd.calculateFood(an)
		if fd.Stock() >= food.amount {
			fd.SetStock(fd.Stock() - food.amount)
			fmt.Printf("Giving %v bracket counts for %s. Left: %v\n", food.amount, an.Type, fd.Stock())
		} else {
			fmt.Println("Not enouch brackets! Need to refill")
			lsChan <- true
			break
		}
	}
}

func (fd *Feeder) calculateFood(an *animal.Animal) FoodBracket {
	amount := 0
	switch an.Type {
	case animal.Bear:
		amount = 6 * an.Weight / 100
	case animal.Deer:
		amount = 4 * an.Weight / 100
	case animal.Lion:
		amount = 4 * an.Weight / 100
	case animal.Wolf:
		amount = 2 * an.Weight / 100
	}
	return FoodBracket{animalType: an.Type, amount: amount}
}

func (fd *Feeder) Refill(amount int) {
	fd.SetStock(fd.Stock() + amount)
	fmt.Printf("Feeder refilled for %v brackets. Left: %v\n", amount, fd.Stock())
}
