package main

type cake struct{}

func (cake) name() string {
	return "cake"
}

func (cake) calories() int {
	return 450
}

type icecream struct{}

func (icecream) name() string {
	return "icecream"
}

func (icecream) calories() int {
	return 200
}

type pizza struct{}

func (pizza) name() string {
	return "pizza"
}

func (pizza) calories() int {
	return 700
}

type butter struct{}

func (butter) name() string {
	return "butter"
}

func (butter) calories() int {
	return 700
}

type bread struct{}

func (bread) name() string {
	return "bread"
}

func (bread) calories() int {
	return 250
}

type cheese struct{}

func (cheese) name() string {
	return "cheese"
}

func (cheese) calories() int {
	return 400
}

type sandwich struct {
	br bread
	bu butter
}

func (sandwich) name() string {
	return "sandwich"
}

func (s *sandwich) calories() int {
	return s.br.calories() + s.bu.calories()
}

type cheeseSandwich struct {
	s  sandwich
	ch cheese
}

func (cheeseSandwich) name() string {
	return "cheese sandwich"
}

func (cs *cheeseSandwich) calories() int {
	return cs.s.calories() + cs.ch.calories()
}

type potato struct{}

func (potato) name() string {
	return "potato"
}

func (potato) calories() int {
	return 75
}
