package camera

import (
	"errors"
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

type PartOfDay string

const (
	Morning = PartOfDay("Morning")
	Day     = PartOfDay("Day")
	Evening = PartOfDay("Evening")
	Night   = PartOfDay("Night")
)

func (pod PartOfDay) String() string {
	return string(pod)
}

var allPartsOfDay = []PartOfDay{Morning, Day, Evening, Night}

func NextPartOfDay(pod PartOfDay) (PartOfDay, error) {
	for i, part := range allPartsOfDay {
		if part == pod {
			if i == len(allPartsOfDay)-1 {
				return allPartsOfDay[0], nil
			}
			return allPartsOfDay[i+1], nil
		}
	}
	return "", errors.New("invalid part of day")
}

type ExternalLightCamera struct {
	id     int
	animal *animal.Animal
}

func NewExternalLight(id int, an *animal.Animal) *ExternalLightCamera {
	return &ExternalLightCamera{id: id, animal: an}
}

func (c *ExternalLightCamera) Process(pod PartOfDay) error {
	if pod == Evening || pod == Night {
		return fmt.Errorf("this external light camera with id %v cannot work at %v", c.id, pod)
	}
	fmt.Printf("Processing external light camera with id = %v\n", c.id)
	return nil
}

type NightLightCamera struct {
	id     int
	animal *animal.Animal
}

func NewNightLight(id int, an *animal.Animal) *NightLightCamera {
	return &NightLightCamera{id: id, animal: an}
}

func (c *NightLightCamera) Process(pod PartOfDay) error {
	if pod == Morning || pod == Day {
		return fmt.Errorf("this night light camera with id %v cannot work at %v", c.id, pod)
	}
	fmt.Printf("Processing night light camera with id = %v\n", c.id)
	return nil
}
