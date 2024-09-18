package animal

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"sync"
)

const (
	MaxIndicatorValue     = 100
	IndicatorCoef         = 30
	CriticalIndicatorCoef = 10
	MaxSatietyCoefDelta   = 40
)

type Animal struct {
	ID int

	indicatorsMu sync.Mutex
	health       int
	mood         int
	satiety      int

	logger *slog.Logger
}

func (an *Animal) String() string {
	return fmt.Sprintf("id = %2v, health = %3v, mood = %3v, satiety = %3v", an.ID, an.Health(), an.Mood(), an.Satiety())
}

func NewAnimal(id int, logger *slog.Logger) *Animal {
	return &Animal{
		ID:      id,
		health:  MaxIndicatorValue,
		mood:    MaxIndicatorValue,
		satiety: MaxIndicatorValue,
		logger:  logger,
	}
}

func (an *Animal) Health() int {
	defer an.indicatorsMu.Unlock()

	an.indicatorsMu.Lock()
	return an.health
}

func (an *Animal) SetHealth(v int) {
	an.indicatorsMu.Lock()
	an.health = v
	an.indicatorsMu.Unlock()
}

func (an *Animal) Mood() int {
	defer an.indicatorsMu.Unlock()

	an.indicatorsMu.Lock()
	return an.mood
}

func (an *Animal) SetMood(v int) {
	an.indicatorsMu.Lock()
	an.mood = v
	an.indicatorsMu.Unlock()
}

func (an *Animal) Satiety() int {
	defer an.indicatorsMu.Unlock()

	an.indicatorsMu.Lock()
	return an.satiety
}

func (an *Animal) SetSatiety(v int) {
	an.indicatorsMu.Lock()
	an.satiety = v
	an.indicatorsMu.Unlock()
}

func (an *Animal) IsHungry() bool {
	return an.Satiety() < IndicatorCoef
}

func (an *Animal) HasCriticalValues() bool {
	return an.Health() <= CriticalIndicatorCoef || an.Satiety() <= CriticalIndicatorCoef || an.Mood() <= CriticalIndicatorCoef
}

func (an *Animal) RandomlyChangeIndicators() {
	an.SetHealth(rand.IntN(MaxIndicatorValue) + 1)
	an.SetMood(rand.IntN(MaxIndicatorValue) + 1)

	//	hunger is more linear
	randDelta := rand.IntN(MaxSatietyCoefDelta-10) + 10
	an.SetSatiety(max(CriticalIndicatorCoef, an.Satiety()-randDelta))

	an.logger.Info("Animal randomly changed it's values: ", "values", an, "is hungry:", an.IsHungry(), "has critical values: ", an.HasCriticalValues())
}

func (an *Animal) eat(food int) {
	an.logger.Info("Animal eats", "id", an.ID, "food amount", food)
	an.SetSatiety(an.Satiety() + food)
}
