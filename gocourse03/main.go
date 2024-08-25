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

type rodentType string

func (r rodentType) String() string {
	return string(r)
}

var rodentTypes = [5]rodentType{"Beaver", "Capybara", "Mouse", "Rat", "Squirrel"}

func generateRandomRodentType() rodentType {
	return rodentTypes[rand.IntN(len(rodentTypes))]
}

type coordinates struct {
	x int
	y int
}

func generateRandomCoordinates() coordinates {
	return coordinates{
		x: rand.IntN(labyrinthSize),
		y: rand.IntN(labyrinthSize),
	}
}

type sector struct {
	leftTop     coordinates
	rightBottom coordinates
	rodents     []*rodent
}

func newSector(lt coordinates, rb coordinates) *sector {
	return &sector{
		leftTop:     lt,
		rightBottom: rb,
		rodents:     []*rodent{},
	}
}

func (s *sector) areCoordinatedInSector(c coordinates) bool {
	return c.x >= s.leftTop.x && c.x <= s.rightBottom.x && c.y >= s.leftTop.y && c.y <= s.rightBottom.y
}

func (s *sector) addRodent(r *rodent) {
	s.rodents = append(s.rodents, r)
}

func (s *sector) removeRodent(r *rodent) {
	for i, rod := range s.rodents {
		if rod.id == r.id {
			s.rodents = append(s.rodents[:i], s.rodents[i+1:]...)
			return
		}
	}
}

func findSector(sectors []*sector, c coordinates) *sector {
	for _, s := range sectors {
		if s.areCoordinatedInSector(c) {
			return s
		}
	}

	fmt.Println("Impossible situation, sector must be found")
	return nil
}

type rodent struct {
	id        int
	tp        rodentType
	movements []coordinates
}

func newRodent(id int) *rodent {
	return &rodent{
		id:        id,
		tp:        generateRandomRodentType(),
		movements: []coordinates{generateRandomCoordinates()},
	}
}

func (r *rodent) String() string {
	return fmt.Sprintf("id = %v, type = %v, movements = %v", r.id, r.tp, r.movements)
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
	}

	fmt.Println("Impossible situation – no valid moves")
	return coordinates{}
}

func main() {
	lCenter := labyrinthSize / 2
	//	meaning that we have 4 same-size sectors
	//	(0, 0) -> (4, 4), (5, 0) -> (9, 4)
	//	(5, 0) -> (5, 4), (5, 5) -> (9, 9)
	sectors := []*sector{
		newSector(coordinates{x: 0, y: 0}, coordinates{x: lCenter - 1, y: lCenter - 1}),
		newSector(coordinates{x: lCenter, y: 0}, coordinates{x: labyrinthSize - 1, y: lCenter - 1}),
		newSector(coordinates{x: 0, y: lCenter - 1}, coordinates{x: lCenter - 1, y: labyrinthSize - 1}),
		newSector(coordinates{x: lCenter, y: lCenter - 1}, coordinates{x: labyrinthSize - 1, y: labyrinthSize - 1}),
	}

	movesCount := 24 * 60
	// at least 1 rodent must be created
	rodentCount := rand.IntN(rodentMaxCount-1) + 1
	fmt.Printf("Create %v rodents\n", rodentCount)
	rodents := make([]*rodent, 0, rodentCount)

	// key – time, value – slice of rodents that were on sensor at this moment
	// actually it can be
	// [time: 3, id: 2], [time: 7, id: 17, 19], [time: 21, id: 2, 3, 5]
	sensorHistory := make(map[int][]*rodent)

	sensorCoords := coordinates{
		x: rand.IntN(labyrinthSize),
		y: rand.IntN(labyrinthSize),
	}

	for i := 0; i < rodentCount; i++ {
		r := newRodent(i + 1)
		rodents = append(rodents, r)
		sector := findSector(sectors, r.movements[0])
		if sector != nil {
			sector.addRodent(r)
		}
	}

	for i := 0; i < movesCount; i++ {
		for _, rodent := range rodents {
			lastCoords := rodent.movements[len(rodent.movements)-1]
			newCoords := makeMove(lastCoords)
			if newCoords == sensorCoords {
				sensorHistory[i] = append(sensorHistory[i], rodent)
			}
			rodent.movements = append(rodent.movements, newCoords)

			currentSector := findSector(sectors, lastCoords)
			newSector := findSector(sectors, newCoords)

			if currentSector != newSector {
				if currentSector != nil {
					currentSector.removeRodent(rodent)
				}
				if newSector != nil {
					newSector.addRodent(rodent)
				}
			}
		}
	}

	fmt.Println("\nMain sensor history:")
	for key, value := range sensorHistory {
		var rodStr string
		for _, rod := range value {
			rodStr += fmt.Sprintf("%v, ", rod.id)
		}
		tm := time.Duration(key) * time.Minute
		fmt.Printf("time = %-10s\trodent ids = %v\n", tm, rodStr)
	}

	fmt.Println("\nInformation about all sectors:")
	for i, sector := range sectors {
		fmt.Printf("Sector %d (%v - %v) contains rodents: ", i+1, sector.leftTop, sector.rightBottom)
		for _, r := range sector.rodents {
			fmt.Printf("%d ", r.id)
		}
		fmt.Println()
	}

	fmt.Println("\nAll rodents and their positions:")
	for _, rodent := range rodents {
		fmt.Printf("Rodend #%-2v, type: %-10v: start position: %v, end position: %v\n", rodent.id, rodent.tp, rodent.movements[0], rodent.movements[len(rodent.movements) - 1])
	}
}
