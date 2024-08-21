package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/vicuani/go_course/gocourse02/internal"
)

/*

Написати програму “Зоопарк”. Звіри повтікали (більше трьох штук), наглядач повинен їх зібрати.
Кожна сутність (наглядач, звір, клітка, …) представляється окремою структурою (zookeeper, Animal, Cage, …).
Треба використати щонайменше: структури, вказівник, nil, вбудовування (embedding), конструктор.
Додати тваринам можливість розмножуватись.
Програма має демонструвати свою роботу через вивід в stdout.

*/

// zookeeper holds animals and cages to always know how many animals and cages there are.
type zookeeper struct {
	animals []*internal.Animal
	cages   []*internal.Cage
}

func newZookeeper() *zookeeper {
	return &zookeeper{}
}

func (k *zookeeper) putToCage(c *internal.Cage, an *internal.Animal) error {
	err := c.PutAnimal(an)
	if err != nil {
		return err
	}

	err = an.SetCage(c)
	if err != nil {
		return err
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
	animalCount := len(k.animals) // pre-stored to give an option to reproduce only for animals that were created before
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
					newC := internal.NewCage()
					k.animals = append(k.animals, newAn)
					k.cages = append(k.cages, newC)
					newId++
				}
			}
		}
	}
}

func printAnimals(animals []*internal.Animal, context string) {
	fmt.Println(context + ":\n")
	for _, an := range animals {
		fmt.Printf("Animal #%v,\tspecies: %v,\tgender: %v,\thas Cage = %v\n", an.Id, an.Species, an.Gender, an.GetCage() != nil)
	}
}

func main() {
	animalCount := rand.IntN(10) + 10

	keeper := newZookeeper()

	for i := 0; i < animalCount; i++ {
		sp := rand.IntN(int(internal.AnimalSpeciesCount))
		an := internal.NewAnimal(i+1, internal.AnimalSpecies(sp))
		c := internal.NewCage()

		err := keeper.putToCage(c, an)
		if err != nil {
			fmt.Println(err)
		} else {
			keeper.animals = append(keeper.animals, an)
			keeper.cages = append(keeper.cages, c)
		}
	}

	printAnimals(keeper.animals, "First stage, everything is created")
	keeper.sleep()
	keeper.runAnimalsTinder()

	fmt.Println("Zoo keeper woke up and catches the animals to put them into the cages")

	for i := 0; i < len(keeper.animals); i++ {
		keeper.putToCage(keeper.cages[i], keeper.animals[i])
	}

	printAnimals(keeper.animals, "Finally")
}
