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
	index := us.sector.GetAnimalIndex(an)
	if index == -1 {
		return fmt.Errorf("such an animal doesn't exist")
	}

	fmt.Printf("Feeding an animal: %v\n", us.sector.animals[index])
	return nil
}
