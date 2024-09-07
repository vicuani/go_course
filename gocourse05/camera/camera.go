package camera

import (
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

type PartOfDay string

const Morning = PartOfDay("Morning")
const Day = PartOfDay("Day")
const Evening = PartOfDay("Evening")
const Night = PartOfDay("Night")

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
	return PartOfDay("non-existable"), fmt.Errorf("invalid part of day")
}

type ExternalLightCamera struct {
	id     int
	animal *animal.Animal
}

func NewExternalLightCamera(id int, an *animal.Animal) *ExternalLightCamera {
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

func NewNightLightCamera(id int, an *animal.Animal) *NightLightCamera {
	return &NightLightCamera{id: id, animal: an}
}

func (c *NightLightCamera) Process(pod PartOfDay) error {
	if pod == Morning || pod == Day {
		return fmt.Errorf("this night light camera with id %v cannot work at %v", c.id, pod)
	}
	fmt.Printf("Processing night light camera with id = %v\n", c.id)
	return nil
}
