package main

import "fmt"

type animal struct {
	id    int
	lChan chan string
}

func newAnimal(id int, lChan chan string) *animal {
	return &animal{id: id, lChan: lChan}
}

func (an *animal) eat(f []food) {
	for _, cf := range f {
		res := fmt.Sprintf("Animal #%v ate: %v, calories: %v", an.id, cf.name(), cf.calories())
		an.lChan <- res
	}
}
