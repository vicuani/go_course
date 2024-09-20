package aquarium

type Aquarium struct {
	size         int
	animal       string
	saltLevel    float64
	contaminants float64
	filterSpeed  float64
	cleaners     float64
}

func (aq *Aquarium) Size() int {
	return aq.size
}

func (aq *Aquarium) Animal() string {
	return aq.animal
}

func (aq *Aquarium) SaltLevel() float64 {
	return aq.saltLevel
}

func (aq *Aquarium) Contaminants() float64 {
	return aq.contaminants
}

func (aq *Aquarium) FilterSpeed() float64 {
	return aq.filterSpeed
}

func (aq *Aquarium) Cleaners() float64 {
	return aq.cleaners
}

func (aq *Aquarium) IncreaseFiltration(coef float64) {
	aq.filterSpeed += coef
}

func (aq *Aquarium) DecreaseFiltration(coef float64) {
	aq.filterSpeed -= coef
}

func (aq *Aquarium) AddSalt(coef float64) {
	aq.saltLevel += coef
}

func (aq *Aquarium) AddCleaners(coef float64) {
	aq.cleaners += coef
}

func (aq *Aquarium) CalculatePollution() float64 {
	return aq.Contaminants() / float64(aq.Size())
}

type AquariumBuilder struct {
	size         int
	animal       string
	saltLevel    float64
	contaminants float64
	filterSpeed  float64
	cleaners     float64
}

func NewBuilder() *AquariumBuilder {
	return &AquariumBuilder{}
}

func (b *AquariumBuilder) SetSize(size int) *AquariumBuilder {
	b.size = size
	return b
}

func (b *AquariumBuilder) SetAnimal(animal string) *AquariumBuilder {
	b.animal = animal
	return b
}

func (b *AquariumBuilder) SetSaltLevel(saltLevel float64) *AquariumBuilder {
	b.saltLevel = saltLevel
	return b
}

func (b *AquariumBuilder) SetContaminants(contaminants float64) *AquariumBuilder {
	b.contaminants = contaminants
	return b
}

func (b *AquariumBuilder) SetFilterSpeed(filterSpeed float64) *AquariumBuilder {
	b.filterSpeed = filterSpeed
	return b
}

func (b *AquariumBuilder) SetCleaners(cleaners float64) *AquariumBuilder {
	b.cleaners = cleaners
	return b
}

func (b *AquariumBuilder) Build() *Aquarium {
	return &Aquarium{
		size:         b.size,
		animal:       b.animal,
		saltLevel:    b.saltLevel,
		contaminants: b.contaminants,
		filterSpeed:  b.filterSpeed,
		cleaners:     b.cleaners,
	}
}
