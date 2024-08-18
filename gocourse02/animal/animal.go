package animal

import (
	"GoCourse/go_course/gocourse02/cage"
	"errors"
	"math/rand/v2"
)

type AnimalSpecies int

const (
	Lion AnimalSpecies = iota
	Tiger
	Wolf
	Fox
	AnimalSpeciesCount // virtual element used to determine 'enum' size
)

func (sp AnimalSpecies) String() string {
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

type gender int

const (
	Male gender = iota
	Female
)

func (g gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	}
	return "NON-BINARY"
}

//	artificial struct, just to use embedded struct
type animalProperties struct {
	Species AnimalSpecies
	Gender  gender
}


type Animal struct {
	Id int
	animalProperties
	cage *cage.Cage
}

func (an *Animal) GetCage() *cage.Cage {
	return an.cage
}

func (an *Animal) SetCage(c *cage.Cage) error {
	if an.cage != nil {
		return errors.New("this animal already has a cage")
	}

	cAn := c.GetAnimal()
	if cAn != nil && cAn != an {
		return errors.New("this cage is already full")
	}

	an.cage = c
	return nil
}

func NewAnimal(id int, species AnimalSpecies, c *cage.Cage) *Animal {
	return &Animal{
		Id: id,
		animalProperties: animalProperties{
			Species: species,
			Gender:  gender(rand.IntN(2)),
		},
		cage: c,
	}
}

func (cur *Animal) Reproduce(other *Animal, newId int) (*Animal, error) {
	if cur.cage != nil || other.cage != nil {
		return nil, errors.New("only free animals can reproduce")
	}

	if (cur.Gender == Male && other.Gender == Female) || (cur.Gender == Female && other.Gender == Male) {
		if cur.Species != other.Species {
			return nil, errors.New("animals with different species cannot reproduce")
		} else {
			return NewAnimal(newId, other.Species, nil), nil
		}
	} else {
		return nil, errors.New("this couple is not reproductable")
	}
}

func (an *Animal) Escape() error {
	if an.cage == nil {
		return errors.New("this animal is already free")
	}

	if an.cage.GetAnimal() != an {
		return errors.New("incorrect data: animal's cage contain no animal or other animal")
	}

	an.cage.FreeAnimal()
	an.cage = nil
	return nil
}
