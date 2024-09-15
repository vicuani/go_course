package feeder

import (
	"log/slog"
	"strconv"
	"sync"

	"github.com/vicuani/go_course/gocourse09/animal"
)

type FoodBracket struct {
	animalType animal.AnimalType
	amount     int
}

type FeederInterface interface {
	Stock() int
	SetStock(int)
	Feed(chan<- bool, []animal.AnimalInterface)
	Refill(int)
}

type Feeder struct {
	stockMu sync.Mutex
	stock   int

	logger *slog.Logger
}

func NewFeeder(initialStock int, logger *slog.Logger) *Feeder {
	return &Feeder{
		stock:  initialStock,
		logger: logger,
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

func (fd *Feeder) Feed(lowStockChan chan<- bool, animals []animal.AnimalInterface) {
	for _, an := range animals {
		food := fd.calculateFood(an)
		if fd.Stock() >= food.amount {
			fd.SetStock(fd.Stock() - food.amount)
			fd.logger.Info("Feeding animal:", "animal type", an.Type(), "brackets count:", strconv.Itoa(food.amount), "left:", strconv.Itoa((fd.Stock())))
		} else {
			fd.logger.Info("Not enough brackets! Need to refill")
			lowStockChan <- true
			break
		}
	}
}

func (fd *Feeder) calculateFood(an animal.AnimalInterface) FoodBracket {
	amount := 0
	switch an.Type() {
	case animal.Bear:
		amount = 6 * an.Weight()
	case animal.Deer:
		amount = 4 * an.Weight()
	case animal.Lion:
		amount = 4 * an.Weight()
	case animal.Wolf:
		amount = 2 * an.Weight()
	}
	amount /= 100
	return FoodBracket{animalType: an.Type(), amount: amount}
}

func (fd *Feeder) Refill(amount int) {
	fd.SetStock(fd.Stock() + amount)
	fd.logger.Info("Feeder refilled for:", "brackets count:", strconv.Itoa(amount), "left:", strconv.Itoa(fd.Stock()))
}
