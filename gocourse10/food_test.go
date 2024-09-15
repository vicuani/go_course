package main

import (
	"testing"
)

func TestCheeseSandwichCalories(t *testing.T) {
	cs := cheeseSandwich{
		s: sandwich{
			br: bread{},
			bu: butter{},
		},
		ch: cheese{},
	}
	const expected = 1350

	got := cs.calories()

	if expected != got {
		t.Errorf("Expected cheese sandwich calories: %d, got: %d", expected, got)
	}
}
