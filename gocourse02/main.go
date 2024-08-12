package main

import (
	"fmt"
	"math/rand/v2"
)

/*

Написати програму “Зоопарк”. Звіри повтікали (більше трьох штук), наглядач повинен їх зібрати.
Кожна сутність (наглядач, звір, клітка, …) представляється окремою структурою (zookeeper, animal, cage, …).
Треба використати щонайменше: структури, вказівник, nil, будування, конструктор.
Додати тваринам можливість розмножуватись.
Програма має демонструвати свою роботу через вивід в stdout.

*/

// Zookeeper
type zookeeper struct {
}

func (zookeeper) put_to_cage(c *cage, an *animal) {
	c.put(an)
}

// Animal

type animal_species int16

const (
	Lion animal_species = iota
	Tiger
	Wolf
	Fox
	MaxVirtualAnimal
)

func (animal_species) to_string(sp animal_species) string {
	switch sp {
	case Lion:
		return "Lion"
	case Tiger:
		return "Tiger"
	case Wolf:
		return "Wolf"
	case Fox:
		return "Fox"
	}
	return "UNKNOWN ANIMAL"
}

type animal struct {
	Id      int
	Species animal_species
	Cage    *cage
}

func (animal) copy(other *animal, new_id int) *animal {
	return &animal{
		Id:      new_id,
		Species: other.Species,
		Cage:    other.Cage,
	}
}

// Cage
type cage struct {
	animals []*animal
}

func (c* cage) free_animals() {
	for _, animal := range c.animals {
		animal.Cage = nil
	}

	c.animals = []*animal{}
}

func (c* cage) put(an *animal) {
	if an.Cage != nil {
		fmt.Printf("This animal (id = %v) is already in cage\n", an.Id)
		return
	}

	an.Cage = c
	c.animals = append(c.animals, an)
}

// main
func multiply_animal(animals *[]*animal, an *animal) {
	new_an := an.copy(an, len(*animals)+1)
	*animals = append(*animals, new_an)
}

func print_animals(animals []*animal) {
	for _, animal := range animals {
		fmt.Printf("Animal #%v, species: %v, has cage = %v\n", animal.Id, animal.Species.to_string(animal.Species), animal.Cage != nil)
	}
}

func main() {
	c := cage{}

	//	Create animals
	animal_count := rand.IntN(5) + 3
	var animals []*animal
	for i := 0; i < animal_count; i++ {
		sp := animal_species(rand.IntN(int(MaxVirtualAnimal)))
		current_an := animal{
			Id:      i + 1,
			Species: sp,
			Cage:    &c,
		}

		animals = append(animals, &current_an)
	}
	c.animals = animals
	fmt.Printf("After creation cage contains %v animals\n\n", len(c.animals))

	c.free_animals()

	fmt.Printf("During freedom cage contains %v animals\n\n", len(c.animals))

	fmt.Println("Free animals:")
	print_animals(animals)

	//	Multiply animals randomly, using len(animals) - to multiply only 'known' animals, multiplied will be added to the end of animals
	free_animal_count := len(animals)
	for i := 0; i < free_animal_count; i++ {
		r := rand.IntN(10)
		if r < 3 {
			fmt.Printf("Lucky day for animal with id = %v\n", animals[i].Id)
			multiply_animal(&animals, animals[i])
		}
	}

	keeper := zookeeper{}
	fmt.Println("\nPut all animals to the cage")
	for _, animal := range animals {
		keeper.put_to_cage(&c, animal)
	}

	//	Print all animals again
	fmt.Println("\nFinally, animals are:")
	print_animals(animals)
}
