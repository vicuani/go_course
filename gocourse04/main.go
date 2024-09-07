/*

При розробці системи «Розумний зоопарк» техлід ще не вирішив, яку базу даних використовувати і для реалізації прототипа вирішили зберігати дані в памʼяті програми, а саме — в мапах.
Треба побудувати складну структуру мап, відобразивши деяку ієрархію частин зоопарку.

Наприклад, зоопарк ділиться на території: копитні, пернаті, примати… та інше.
Кожна територія має декілька загонів для кожного виду тварин.
Кожен загін може мати ділянку для перебування тварини і технічне приміщення, де можуть зберігатися приладдя для роботи з тваринами.
На кожній території є певна кількість тварин, а технічне приміщення «здатне» виконувати деякі функції, такі як прибирання і годування тварин.
Реалізувати функції пошуку тварини за імʼям або ID, переміщення тварини із загону в загін, годування тварини.

*/

package main

import (
	"fmt"
	"math/rand/v2"
)

type Animal struct {
	id         int
	name       string
	animalType string
}

func NewAnimal(id int, animalType string, name string) *Animal {
	return &Animal{
		id:         id,
		animalType: animalType,
		name:       name,
	}
}

type Area struct {
	animalType string
	sectors    []*Sector
}

type Areas map[int]*Area

func NewArea(animalType string) *Area {
	return &Area{
		animalType: animalType,
		sectors:    []*Sector{},
	}
}

type Zoo struct {
	areas Areas
}

func main() {
	zoo := Zoo{
		areas: Areas{
			1: NewArea("fishes"),
			2: NewArea("amphibians"),
			3: NewArea("reptiles"),
			4: NewArea("birds"),
			5: NewArea("mammals"),
		},
	}

	for _, area := range zoo.areas {
		sectorsCount := rand.IntN(5) + 2
		for i := 0; i < sectorsCount; i++ {
			sector := NewSector(area)
			animalsCount := rand.IntN(10) + 1
			for j := 0; j < animalsCount; j++ {
				sector.AddAnimal(sector.NewRandomAnimal())
			}

			area.sectors = append(area.sectors, sector)
		}
	}

	birdsArea := zoo.areas[4]
	firstBirdsSector := birdsArea.sectors[0]
	firstBirdsSector.utilitySpace.Clean()
	randBird := firstBirdsSector.animals[rand.IntN(len(firstBirdsSector.animals))]
	firstBirdsSector.utilitySpace.Feed(randBird)
	secondBirdsSector := birdsArea.sectors[1]
	err := firstBirdsSector.MoveAnimalToOtherSector(secondBirdsSector, randBird)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")
}
