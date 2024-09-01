package animal

import (
	"fmt"
	"math/rand/v2"
)

const MaxIndicatorValue = 100
const IndicatorCoef = 30
const CriticalIndicatorCoef = 10

const MaxSatietyCoefDelta = 40

type Animal struct {
	ID      int
	Health  float64
	Mood    float64
	Satiety float64
	busy    bool
}

func NewAnimal(id int) *Animal {
	return &Animal{
		ID:      id,
		Health:  MaxIndicatorValue,
		Mood:    MaxIndicatorValue,
		Satiety: MaxIndicatorValue,
		busy:    false,
	}
}

func GenerateAnimals(n int) []*Animal {
	var animals []*Animal
	for i := 0; i < n; i++ {
		animal := NewAnimal(i)
		animals = append(animals, animal)
	}
	return animals
}

func (an *Animal) HasCriticalValues() bool {
	return an.Health < CriticalIndicatorCoef || an.Satiety < CriticalIndicatorCoef || an.Mood < CriticalIndicatorCoef
}

func (an *Animal) IsHungry() bool {
	return an.Satiety < IndicatorCoef
}

func (an *Animal) IsBusy() bool {
	return an.busy
}

func (an *Animal) free() {
	an.busy = false
}

func (an *Animal) RandomlyChangeIndicators() {
	if an.IsBusy() {
		fmt.Printf("Cannot change indicators for animal #%v, it is busy\n", an.ID)
		return
	}

	defer an.free()

	an.busy = true
	an.Health = float64(rand.IntN(MaxIndicatorValue) + 1)
	an.Mood = float64(rand.IntN(MaxIndicatorValue) + 1)

	//	hunger is more linear
	randDelta := float64(rand.IntN(int(MaxSatietyCoefDelta-10) + 10))
	an.Satiety = max(CriticalIndicatorCoef, an.Satiety-randDelta)

	fmt.Printf("Animal randomly changed it's values: %v, has critical values: %v\n", *an, an.HasCriticalValues())
}

func (an *Animal) eat(f *Feeder) error {
	if an.IsBusy() {
		return fmt.Errorf("this animal #%v cannot eat, it is busy", an.ID)
	}

	if an.Satiety > IndicatorCoef {
		return fmt.Errorf("this animal #%v is not yet hungry", an.ID)
	}

	defer an.free()

	an.busy = true
	eaten := min(MaxIndicatorValue-an.Satiety, f.volume)
	an.Satiety += eaten
	f.volume -= eaten
	fmt.Printf("Animal #%v ate %v food from feeder #%v. Now it's satiety = %v, feeders volume = %v\n", an.ID, eaten, f.ID, an.Satiety, f.volume)
	return nil
}
