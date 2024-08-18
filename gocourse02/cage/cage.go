package cage

import (
	"errors"
)

type IAnimal interface {
	GetCage() *Cage
	SetCage(c *Cage) error
}

type Cage struct {
	animal IAnimal
}

func NewCage() *Cage {
	return &Cage{
		animal: nil,
	}
}

func (c *Cage) GetAnimal() IAnimal {
	return c.animal
}

func (c *Cage) PutAnimal(an IAnimal) error {
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
