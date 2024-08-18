package main

import (
	"GoCourse/go_course/gocourse02/animal"
	"GoCourse/go_course/gocourse02/cage"
	"fmt"
	"math/rand/v2"
)

/*

Написати програму “Зоопарк”. Звіри повтікали (більше трьох штук), наглядач повинен їх зібрати.
Кожна сутність (наглядач, звір, клітка, …) представляється окремою структурою (zookeeper, Animal, Cage, …).
Треба використати щонайменше: структури, вказівник, nil, вбудовування (embedding), конструктор.
Додати тваринам можливість розмножуватись.
Програма має демонструвати свою роботу через вивід в stdout.

*/

//	NOTES to implementation
//	Moving animal and cage to separate packages were done for encapsulation (correct using pointers to each other) and readability

type zookeeper struct {
	//	NOTE animals and cages are stored here for the purpose zoo keeper always to know how many animals and cages they have
	animals []*animal.Animal
	cages   []*cage.Cage
}

func newZookeeper() *zookeeper {
	return &zookeeper{}
}

func (k *zookeeper) putToCage(c *cage.Cage, an *animal.Animal, areNew bool) error {
	err := c.PutAnimal(an)
	if err != nil {
		return err
	}

	err = an.SetCage(c)
	if err != nil {
		return err
	}

	if areNew {
		k.animals = append(k.animals, an)
		k.cages = append(k.cages, c)
	}
	return nil
}

func (k *zookeeper) sleep() {
	fmt.Println("zzzzz")
	for _, an := range k.animals {
		err := an.Escape()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (k *zookeeper) runAnimalsTinder() {
	// NOTE this variable is pre-stored to give an option to reproduce only for animals that were created before
	animalCount := len(k.animals)
	newId := animalCount + 1
	for i := 0; i < animalCount; i++ {
		for j := i + 1; j < animalCount; j++ {
			if k.animals[i].Species == k.animals[j].Species && k.animals[i].Gender != k.animals[j].Gender {
				newAn, err := k.animals[i].Reproduce(k.animals[j], newId)
				if err != nil {
					fmt.Println(err, k.animals[i], k.animals[j])
					continue
				} else {
					fmt.Println("We have a match!")
					newC := cage.NewCage()
					k.animals = append(k.animals, newAn)
					k.cages = append(k.cages, newC)
					newId++
				}
			}
		}
	}
}

func printAnimals(animals []*animal.Animal, context string) {
	fmt.Println(context + ":\n")
	for _, an := range animals {
		fmt.Printf("Animal #%v,\tspecies: %v,\tgender: %v,\thas Cage = %v\n", an.Id, an.Species, an.Gender, an.GetCage() != nil)
	}
}

func main() {
	animalCount := rand.IntN(10) + 10

	keeper := newZookeeper()

	for i := 0; i < animalCount; i++ {
		sp := rand.IntN(int(animal.AnimalSpeciesCount))
		an := animal.NewAnimal(i+1, animal.AnimalSpecies(sp), nil)
		c := cage.NewCage()

		err := keeper.putToCage(c, an, true)
		if err != nil {
			fmt.Println(err)
		}
	}

	printAnimals(keeper.animals, "First stage, everything is created")
	keeper.sleep()
	keeper.runAnimalsTinder()

	fmt.Println("Zoo keeper woke up and catches the animals to put them into the cages")

	for i := 0; i < len(keeper.animals); i++ {
		keeper.putToCage(keeper.cages[i], keeper.animals[i], false)
	}

	printAnimals(keeper.animals, "Finally")
}
