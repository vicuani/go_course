package main

import "fmt"

type animal struct {
	id      int
	logChan chan string
}

func newAnimal(id int, logChan chan string) *animal {
	return &animal{id: id, logChan: logChan}
}

func (an *animal) eat(f []food) {
	for _, cf := range f {
		res := fmt.Sprintf("Animal #%v ate: %v, calories: %v", an.id, cf.name(), cf.calories())
		an.logChan <- res
	}
}
