package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

/*

2. Написати програму «Мегаполіс мишей».

У зоопарку є окрема кімната «малі гризуни», де знаходиться кілька сімейств кількох видів гризунів. У кожного є чіп з його ID і видом. Для них у цій кімнаті побудована мережа лабіринтів.
У цьому лабіринті є один датчик, через який впродовж дня проходять сотні гризунів, а до ночі вони розходяться по своїх куточках.

Для працівників зоопарку треба зробити програму, яка буде давати звіт, де знаходився кожний гризун на початку дня і де зупинився наприкінці, а також зберігати історію його рухів через датчик.

У програмі може бути тип/структура для певного виду гризуна, набори (слайси) цих гризунів, і слайс для зберігання руху гризунів. Також стан сектору лабіринту: скільки і які гризуни знаходиться там у певний момент.
Треба симулювати рух гризунів через головний датчик і датчики секторів. Виконати операції видалення, пошуку і додавання в слайс.

Базові типи можуть мати вигляд:

type (
	Sector     string
	rodentType string
	FromTo     [2]Sector

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID      int
		Type    rodentType
		History FromTo
	}
)

*/

const labyrinthSize = 10
const rodentMaxCount = 20

// rodentType
type rodentType int

const (
	Beaver rodentType = iota
	Capybara
	Mouse
	Rat
	Squirrel
	RodentTypesCount // virtual element used to determine 'enum' size
)

func (r rodentType) String() string {
	switch r {
	case Beaver:
		return "Beaver"
	case Capybara:
		return "Capybara"
	case Mouse:
		return "Mouse"
	case Rat:
		return "Rat"
	case Squirrel:
		return "Squirrel"
	}
	return "UNKNOWN"
}

type coordinates struct {
	x int
	y int
}

type rodent struct {
	id          int
	tp          rodentType
	startCoords coordinates
	endCoords   coordinates
	//	TODO add movements
}

func newRodent(id int, tp rodentType, coords coordinates) *rodent {
	return &rodent{
		id:          id,
		tp:          tp,
		startCoords: coords,
		endCoords:   coords,
	}
}

func (r *rodent) String() string {
	return fmt.Sprintf("id = %v, type = %v, start coords = %v, end coords = %v", r.id, r.tp.String(), r.startCoords, r.endCoords)
}

var allMoves = []coordinates{
	{-1, 0}, // Left
	{1, 0},  // Right
	{0, -1}, // Up
	{0, 1},  // Down
}

func makeMove(coords coordinates) coordinates {
	isPossible := func(x, y int) bool {
		return x >= 0 && x < labyrinthSize && y >= 0 && y < labyrinthSize
	}

	validMoves := make([]coordinates, len(allMoves))
	copy(validMoves, allMoves)

	for i := len(validMoves) - 1; i >= 0; i-- {
		move := validMoves[i]
		newX, newY := coords.x+move.x, coords.y+move.y
		if !isPossible(newX, newY) {
			validMoves = append(validMoves[:i], validMoves[i+1:]...)
		}
	}

	if len(validMoves) > 0 {
		choice := rand.IntN(len(validMoves))
		move := validMoves[choice]
		return coordinates{coords.x + move.x, coords.y + move.y}
	} else {
		fmt.Println("Impossible situation – no valid moves")
		return coordinates{}
	}
}

func main() {
	movesCount := 24 * 60
	//	NOTE at least 1 rodent must be created
	rodentCount := rand.IntN(rodentMaxCount-1) + 1
	fmt.Printf("Create %v rodents\n", rodentCount)
	rodents := make([]*rodent, 0, rodentCount)

	//	NOTE key – time, value – slice of rodents that were on sensor at this moment
	// actually it can be
	// [time: 3, id: 2], [time: 7, id: 17, 19], [time: 35, id: 2, 3, 5]
	sensorHistory := make(map[int][]*rodent)

	sensorCoords := coordinates{
		x: rand.IntN(labyrinthSize),
		y: rand.IntN(labyrinthSize),
	}

	for i := 0; i < rodentCount; i++ {
		tp := rodentType(rand.IntN(int(RodentTypesCount)))
		coords := coordinates{
			x: rand.IntN(labyrinthSize),
			y: rand.IntN(labyrinthSize),
		}

		r := newRodent(i+1, tp, coords)
		rodents = append(rodents, r)
	}

	for i := 0; i < movesCount; i++ {
		for _, rodent := range rodents {
			newCoords := makeMove(rodent.endCoords)
			if newCoords == sensorCoords {
				sensorHistory[i] = append(sensorHistory[i], rodent)
			}
			rodent.endCoords = newCoords
		}
	}

	for key, value := range sensorHistory {
		var rodStr string
		for j, rod := range value {
			rodStr += fmt.Sprintf("%v : %v; ", j, rod.String())
		}
		tm := time.Duration(key) * time.Minute
		fmt.Printf("time = %s,\t\tvalue = %v\n", tm, value)
	}

	//	TODO print all rodent + it's start position
}
