package animal

import "sync"

type AnimalType string

const (
	Bear AnimalType = "Bear"
	Deer AnimalType = "Deer"
	Lion AnimalType = "Lion"
	Wolf AnimalType = "Wolf"
)

type Animal struct {
	ID     int
	Type   AnimalType
	Weight int

	inZoneMu sync.Mutex
	inZone   bool
}

func NewAnimal(anType AnimalType, weight int) *Animal {
	return &Animal{
		Type:   anType,
		Weight: weight,
	}
}

func GenerateAnimals() []*Animal {
	return []*Animal{
		NewAnimal(Bear, 200),
		NewAnimal(Deer, 120),
		NewAnimal(Lion, 150),
		NewAnimal(Wolf, 50),
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

type AnimalDetector struct{}

func (ad *AnimalDetector) Detect(zone *Zone) []*Animal {
	zone.mu.Lock()
	defer zone.mu.Unlock()

	var detectedAnimals []*Animal
	for _, animal := range zone.Animals {
		if animal.InZone() {
			detectedAnimals = append(detectedAnimals, animal)
		}
	}
	return detectedAnimals
}

type Zone struct {
	Animals []*Animal
	mu      sync.Mutex
}

func GenerateZone() *Zone {
	return &Zone{Animals: GenerateAnimals()}
}
