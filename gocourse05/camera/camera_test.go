package camera

import (
	"testing"

	"github.com/vicuani/go_course/gocourse05/animal"
)

func preconditionsExternalLight() *ExternalLightCamera {
	id := 4
	animal := animal.NewAnimal(id)
	return NewExternalLightCamera(id, animal)
}

func TestExternalLightCamera(t *testing.T) {
	camera := preconditionsExternalLight()

	positiveCheck := func(pod PartOfDay) {
		err := camera.Process(pod)
		if err != nil {
			t.Errorf("External light camera doesn't work in the %v despite it should\n", pod)
		}
	}

	negativeCheck := func(pod PartOfDay) {
		err := camera.Process(pod)
		if err == nil {
			t.Errorf("External light camera works in the %v despite it shouldn't\n", pod)
		}
	}

	positiveCheck(PartOfDay("Morning"))
	positiveCheck(PartOfDay("Day"))
	negativeCheck(PartOfDay("Evening"))
	negativeCheck(PartOfDay("Night"))
}

func preconditionsNightLight() *NightLightCamera {
	id := 14
	animal := animal.NewAnimal(id)
	return NewNightLightCamera(id, animal)
}

func TestNightLightCamera(t *testing.T) {
	camera := preconditionsNightLight()

	positiveCheck := func(pod PartOfDay) {
		err := camera.Process(pod)
		if err != nil {
			t.Errorf("Night light camera doesn't work in the %v despite it should\n", pod)
		}
	}

	negativeCheck := func(pod PartOfDay) {
		err := camera.Process(pod)
		if err == nil {
			t.Errorf("Night light camera works in the %v despite it shouldn't\n", pod)
		}
	}

	negativeCheck(PartOfDay("Morning"))
	negativeCheck(PartOfDay("Day"))
	positiveCheck(PartOfDay("Evening"))
	positiveCheck(PartOfDay("Night"))
}
