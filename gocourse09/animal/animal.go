package animal

import "sync"

type AnimalType string

const (
	Bear AnimalType = "Bear"
	Deer AnimalType = "Deer"
	Lion AnimalType = "Lion"
	Wolf AnimalType = "Wolf"
)

type AnimalInterface interface {
	InZone() bool
	SetInZone(bool)
	Type() AnimalType
	Weight() int
}

type Animal struct {
	ID     int
	anType AnimalType
	weight int

	inZoneMu sync.Mutex
	inZone   bool
}

func NewAnimal(anType AnimalType, weight int) *Animal {
	return &Animal{
		anType: anType,
		weight: weight,
	}
}

func (an *Animal) InZone() bool {
	defer an.inZoneMu.Unlock()

	an.inZoneMu.Lock()
	return an.inZone
}

func (an *Animal) SetInZone(v bool) {
	an.inZoneMu.Lock()
	an.inZone = v
	an.inZoneMu.Unlock()
}

func (an *Animal) Type() AnimalType {
	return an.anType
}

func (an *Animal) Weight() int {
	return an.weight
}

type Detector struct{}

func (ad *Detector) Detect(zone *Zone) []AnimalInterface {
	zone.animalsMu.Lock()
	defer zone.animalsMu.Unlock()

	var detectedAnimals []AnimalInterface
	for _, animal := range zone.Animals {
		if animal.InZone() {
			detectedAnimals = append(detectedAnimals, animal)
		}
	}
	return detectedAnimals
}

type Zone struct {
	Animals   []AnimalInterface
	animalsMu sync.Mutex
}
