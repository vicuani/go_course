package main

type feeder struct {
	ordinaryFood []food
	weekendFood  []food

	eaters []eater
	fs     feedingStrategy
}

func (cs *feeder) addEater(o eater) {
	cs.eaters = append(cs.eaters, o)
}

func (cs *feeder) feedAll() {
	for _, eater := range cs.eaters {
		eater.eat(cs.fs.getFood())
	}
}

func (cs *feeder) strategy() feedingStrategy {
	return cs.fs
}

func (cs *feeder) setStrategy(fs feedingStrategy) {
	cs.fs = fs
}
