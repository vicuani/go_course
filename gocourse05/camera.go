package main

import "fmt"

type partOfDay string

func (pod partOfDay) String() string {
	return string(pod)
}

var allPartsOfDay = []partOfDay{"Morning", "Day", "Evening", "Night"}

func nextPartOfDay(pod partOfDay) partOfDay {
	for i, part := range allPartsOfDay {
		if part == pod {
			if i == len(allPartsOfDay)-1 {
				return allPartsOfDay[0]
			}
			return allPartsOfDay[i+1]
		}
	}
	return "INVALID"
}

type camera interface {
	Process(pod partOfDay) error
}

type externalLightCamera struct {
	id     int
	animal *animal
}

func newExternalLightCamera(id int, an *animal) *externalLightCamera {
	return &externalLightCamera{id: id, animal: an}
}

func (elc *externalLightCamera) Process(pod partOfDay) error {
	if pod == "Evening" || pod == "Night" {
		return fmt.Errorf("this camera cannot work at %v", pod)
	}
	fmt.Printf("Processing external light camera with id = %v", elc.id)
	return nil
}

type nightLightCamera struct {
	id     int
	animal *animal
}

func newNightLightCamera(id int, an *animal) *nightLightCamera {
	return &nightLightCamera{id: id, animal: an}
}

func (nlc *nightLightCamera) Process(pod partOfDay) error {
	if pod == "Morning" || pod == "Day" {
		return fmt.Errorf("this camera cannot work at %v", pod)
	}
	fmt.Printf("Processing night light camera with id = %v", nlc.id)
	return nil
}
