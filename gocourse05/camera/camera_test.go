package camera

import (
	"testing"

	"github.com/vicuani/go_course/gocourse05/animal"
)

func TestExternalLightCamera_Process(t *testing.T) {
	t.Run("when Morning should return nil", func(t *testing.T) {
		pod := PartOfDay("Morning")
		animal := animal.NewAnimal(4)
		camera := NewExternalLight(4, animal)

		err := camera.Process(pod)
		if err != nil {
			t.Errorf("Process() failed: got=%v, but want=<nil>", err)
		}
	})
	t.Run("when Evening should return error", func(t *testing.T) {
		pod := PartOfDay("Evening")
		animal := animal.NewAnimal(7)
		camera := NewExternalLight(7, animal)

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
		camera := NewNightLight(6, animal)

		err := camera.Process(pod)
		if err != nil {
			t.Errorf("Process() failed: got=%v, but want=<nil>", err)
		}
	})
	t.Run("when Day should return error", func(t *testing.T) {
		pod := PartOfDay("Day")
		animal := animal.NewAnimal(17)
		camera := NewNightLight(17, animal)

		err := camera.Process(pod)

		if err == nil {
			t.Error("Process() failed: got=<nil>, but want error")
		}
	})
}

func TestNextPartOfDay(t *testing.T) {
	t.Run("positive1", func(t *testing.T) {
		res, err := NextPartOfDay(Morning)
		if err != nil {
			t.Errorf("error must be nil: %v\n", err)
		}

		if res != Day {
			t.Errorf("incorrect next part of day: %v\n", res)
		}
	})

	t.Run("positive2", func(t *testing.T) {
		res, err := NextPartOfDay(Night)
		if err != nil {
			t.Errorf("error must be nil: %v\n", err)
		}

		if res != Morning {
			t.Errorf("incorrect next part of day: %v\n", res)
		}
	})

	t.Run("negative", func(t *testing.T) {
		res, err := NextPartOfDay("Branch")
		if err == nil {
			t.Errorf("error must not be nil: %v\n", err)
		}

		if res != "" {
			t.Errorf("incorrect empty next part of day: %v\n", res)
		}
	})
}
