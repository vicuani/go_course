package internal

import (
	"errors"
)

type Cage struct {
	animal *Animal
}

func NewCage() *Cage {
	return &Cage{
		animal: nil,
	}
}

func (c *Cage) GetAnimal() *Animal {
	return c.animal
}

func (c *Cage) PutAnimal(an *Animal) error {
	if c.animal != nil {
		return errors.New("this cage is already full")
	}

	anC := an.GetCage()
	if anC != nil && anC != c {
		return errors.New("this animal already has a cage")
	}

	c.animal = an
	return nil
}

func (c *Cage) FreeAnimal() {
	c.animal = nil
}
