package main

import (
	"testing"
)

func TestSandwichCalories(t *testing.T) {
	br := bread{}
	bu := butter{}
	s := sandwich{br: br, bu: bu}

	expectedCalories := br.calories() + bu.calories()
	if s.calories() != expectedCalories {
		t.Errorf("Expected sandwich calories: %d, got: %d", expectedCalories, s.calories())
	}

	ch := cheese{}
	cs := cheeseSandwich{s: s, ch: ch}

	expectedCheeseSandwichCalories := s.calories() + ch.calories()
	if cs.calories() != expectedCheeseSandwichCalories {
		t.Errorf("Expected cheese sandwich calories: %d, got: %d", expectedCheeseSandwichCalories, cs.calories())
	}
}
