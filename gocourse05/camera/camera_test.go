package camera

import (
	"testing"

	"github.com/vicuani/go_course/gocourse05/animal"
)

func TestExternalLightCamera_Process(t *testing.T) {
	t.Run("when Morning should return nil", func(t *testing.T) {
		pod := PartOfDay("Morning")
		animal := animal.NewAnimal(4)
		camera := NewExternalLightCamera(4, animal)

		err := camera.Process(pod)

		if err != nil {
			t.Errorf("Process() failed: got=%v, but want=<nil>", err)
		}
	})
	t.Run("when Evening should return error", func(t *testing.T) {
		pod := PartOfDay("Evening")
		animal := animal.NewAnimal(7)
		camera := NewExternalLightCamera(7, animal)

		err := camera.Process(pod)

		if err == nil {
			t.Error("Process() failed: got=<nil>, but want error")
		}
	})
}

func TestNightlLightCamera_Process(t *testing.T) {
	t.Run("when Night should return nil", func(t *testing.T) {
		pod := PartOfDay("Night")
		animal := animal.NewAnimal(6)
		camera := NewNightLightCamera(6, animal)

		err := camera.Process(pod)

		if err != nil {
			t.Errorf("Process() failed: got=%v, but want=<nil>", err)
		}
	})
	t.Run("when Day should return error", func(t *testing.T) {
		pod := PartOfDay("Day")
		animal := animal.NewAnimal(17)
		camera := NewNightLightCamera(17, animal)

		err := camera.Process(pod)

		if err == nil {
			t.Error("Process() failed: got=<nil>, but want error")
		}
	})
}
