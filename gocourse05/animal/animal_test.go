package animal

import (
	"testing"
)

func TestIsAnimalStateDangerous(t *testing.T) {
	an := NewAnimal(4)

	t.Run("negative1", func(t *testing.T) {
		an.state = AnimalState("Running")
		if an.IsAnimalStateDangerous() {
			t.Error("State 'Running' is not dangerous")
		}
	})

	t.Run("negative2", func(t *testing.T) {
		an.state = AnimalState("Eating")
		if an.IsAnimalStateDangerous() {
			t.Error("State 'Eating' is not dangerous")
		}
	})

	t.Run("positive1", func(t *testing.T) {
		an.state = AnimalState("Fighting")
		if !an.IsAnimalStateDangerous() {
			t.Error("State 'Fighting' is dangerous")
		}
	})

	t.Run("positive2", func(t *testing.T) {
		an.state = AnimalState("Escaping")
		if !an.IsAnimalStateDangerous() {
			t.Error("State 'Escaping' is dangerous")
		}
	})
}
