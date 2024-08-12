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

func (zookeeper) putToCage(c *cage, an *animal) {
	c.put(an)
}

// Animal

type animalSpecies int

const (
	Lion animalSpecies = iota
	Tiger
	Wolf
	Fox
	AnimalCount	//	virtual element used to determine 'enum' size
)

func (sp animalSpecies) String() string {
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
	Species animalSpecies
	Cage    *cage
}

func newAnimal(id int, species animalSpecies, c *cage) *animal {
	return &animal{
		Id: id,
		Species: species,
		Cage: c,
	}
}

func (*animal) reproduce(other *animal, new_id int) *animal {
	return &animal{
		Id:      new_id,
		Species: other.Species,
		Cage:    other.Cage,
	}
}

func (an *animal) excape() {
	an.Cage = nil
}

// Cage
type cage struct {
	animals []*animal
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
func reproduceAnimal(animals *[]*animal, an *animal) {
	new_an := an.reproduce(an, len(*animals)+1)
	*animals = append(*animals, new_an)
}

func printAnimals(animals []*animal) {
	for _, animal := range animals {
		fmt.Printf("Animal #%v, species: %v, has cage = %v\n", animal.Id, animal.Species, animal.Cage != nil)
	}
}

func main() {
	c := cage{}

	//	Create animals
	animal_count := rand.IntN(5) + 3
	var animals []*animal
	for i := 0; i < animal_count; i++ {
		sp := animalSpecies(rand.IntN(int(AnimalCount)))
		animals = append(animals, newAnimal(i + 1, sp, &c))
	}
	c.animals = animals
	fmt.Printf("After creation cage contains %v animals\n\n", len(c.animals))

	for _, animal := range(c.animals) {
		animal.excape()
	}
	c.animals = nil

	fmt.Printf("During freedom cage contains %v animals\n\n", len(c.animals))

	fmt.Println("Free animals:")
	printAnimals(animals)

	//	Multiply animals randomly, using len(animals) - to multiply only 'known' animals, multiplied will be added to the end of animals
	free_animal_count := len(animals)
	for i := 0; i < free_animal_count; i++ {
		r := rand.IntN(10)
		if r < 3 {
			fmt.Printf("Lucky day for animal with id = %v\n", animals[i].Id)
			reproduceAnimal(&animals, animals[i])
		}
	}

	keeper := zookeeper{}
	fmt.Println("\nPut all animals to the cage")
	for _, animal := range animals {
		keeper.putToCage(&c, animal)
	}

	//	Print all animals again
	fmt.Println("\nFinally, animals are:")
	printAnimals(animals)
}
