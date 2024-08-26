package main

import (
	"fmt"
	"math/rand/v2"
)

/*
3. Написати програму «Зооспостереження»

Треба зробити бекенд для сервера, на який надходять нічні зображення з різних камер спостереження з датчиком руху, розвішаних по всьому зоопарку.
У зоопарку існує кілька типів камер. Деякі камери працюють із зовнішнім світлом, інші — в нічному режимі.
Треба обробляти дані з різних джерел (типів камер), зберігати в памʼяті історію подій і передавати єдиний уніфікований запит на інший сервер.
Відповідно, треба зробити кілька типів (структур) які відповідають своїм реальним камерам, і декілька інтерфейсів, із якими працює програма. Використовувати контракти і обробляти можливі помилки.

Також треба написати тести для позитивних і негативних випадків роботи функцій, які оброблюють дані з камер.

Тут «сервер» — умовна назва для нашої програми.

Даними може виступати рух певної тварини. Наприклад: тигр, пішов ліворуч; ведмідь, стоїть.

*/

func createAnimalAndCamera(id int) (camera, *animal) {
	var cam camera
	an := newAnimal(id)
	if rand.IntN(2) == 0 {
		cam = newExternalLightCamera(id, an)
	} else {
		cam = newNightLightCamera(id, an)
	}
	return cam, an
}

func main() {
	cameraCount := rand.IntN(10) + 5

	var cameras []camera
	var animals []*animal

	for i := range cameraCount {
		cam, an := createAnimalAndCamera(i)
		cameras = append(cameras, cam)
		animals = append(animals, an)
	}

	history := animalHistory{}

	movesCount := 10
	timeOfDay := partOfDay("Morning")
	for i := 0; i < movesCount; i++ {
		timeOfDay = nextPartOfDay(timeOfDay)
		currentHistory := []*animal{}
		for j := 0; j < cameraCount; j++ {
			animals[j].state = generateRandomAnimalState()
			err := cameras[j].Process(timeOfDay)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if isAnimalStateDangerous(animals[j].state) {
				fmt.Printf("Animal with id %v has dangerous state: %v", animals[j].id, animals[j].state)
			}
			currentHistory = append(currentHistory, animals[j])
		}
		history[i] = currentHistory
	}

	fmt.Println("\nAll history:")
	fmt.Println(history)
}
