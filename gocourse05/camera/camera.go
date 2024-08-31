package camera

import (
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

type PartOfDay string

func (pod PartOfDay) String() string {
	return string(pod)
}

var allPartsOfDay = []PartOfDay{"Morning", "Day", "Evening", "Night"}

func NextPartOfDay(pod PartOfDay) PartOfDay {
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

type Camera interface {
	Process(pod PartOfDay) error
}

type ExternalLightCamera struct {
	id     int
	animal *animal.Animal
}

func NewExternalLightCamera(id int, an *animal.Animal) *ExternalLightCamera {
	return &ExternalLightCamera{id: id, animal: an}
}

func (elc *ExternalLightCamera) Process(pod PartOfDay) error {
	if pod == "Evening" || pod == "Night" {
		return fmt.Errorf("this external light camera with id %v cannot work at %v", elc.id, pod)
	}
	fmt.Printf("Processing external light camera with id = %v\n", elc.id)
	return nil
}

type NightLightCamera struct {
	id     int
	animal *animal.Animal
}

func NewNightLightCamera(id int, an *animal.Animal) *NightLightCamera {
	return &NightLightCamera{id: id, animal: an}
}

func (nlc *NightLightCamera) Process(pod PartOfDay) error {
	if pod == "Morning" || pod == "Day" {
		return fmt.Errorf("this night light camera with id %v cannot work at %v", nlc.id, pod)
	}
	fmt.Printf("Processing night light camera with id = %v\n", nlc.id)
	return nil
}
