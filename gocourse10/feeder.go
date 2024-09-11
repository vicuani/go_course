package main

type feeder struct {
	ordinaryFood []food
	weekendFood  []food

	observers []observer
	fs        feedingStrategy
}

func (cs *feeder) addObserver(o observer) {
	cs.observers = append(cs.observers, o)
}

func (cs *feeder) feedAll() {
	for _, observer := range cs.observers {
		observer.eat(cs.fs.getFood())
	}
}

func (cs *feeder) strategy() feedingStrategy {
	return cs.fs
}

func (cs *feeder) setStrategy(fs feedingStrategy) {
	cs.fs = fs
}
