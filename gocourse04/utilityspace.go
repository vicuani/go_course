package main

import "fmt"

type UtilitySpace struct {
	sector *Sector
}

func NewUtilitySpace(sector *Sector) *UtilitySpace {
	return &UtilitySpace{
		sector: sector,
	}
}

func (us *UtilitySpace) Clean() {
	fmt.Println("Fröken Bock is on hers way")
}

func (us *UtilitySpace) Feed(an *Animal) error {
	contains := us.sector.ContainsAnimal(an)
	if !contains {
		return fmt.Errorf("such an animal doesn't exist")
	}

	fmt.Println("Feeding an animal")
	return nil
}
