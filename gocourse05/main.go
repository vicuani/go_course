/*
Написати програму «Зооспостереження»

Треба зробити бекенд для сервера, на який надходять нічні зображення з різних камер спостереження з датчиком руху, розвішаних по всьому зоопарку.
У зоопарку існує кілька типів камер. Деякі камери працюють із зовнішнім світлом, інші — в нічному режимі.
Треба обробляти дані з різних джерел (типів камер), зберігати в памʼяті історію подій і передавати єдиний уніфікований запит на інший сервер.
Відповідно, треба зробити кілька типів (структур) які відповідають своїм реальним камерам, і декілька інтерфейсів, із якими працює програма. Використовувати контракти і обробляти можливі помилки.

Також треба написати тести для позитивних і негативних випадків роботи функцій, які оброблюють дані з камер.

Тут «сервер» — умовна назва для нашої програми.

Даними може виступати рух певної тварини. Наприклад: тигр, пішов ліворуч; ведмідь, стоїть.

*/

package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/vicuani/go_course/gocourse05/animal"
	"github.com/vicuani/go_course/gocourse05/camera"
	"github.com/vicuani/go_course/gocourse05/server"
)

func createAnimalAndCamera(id int) (camera.Camera, *animal.Animal) {
	var cam camera.Camera
	an := animal.NewAnimal(id)
	if rand.IntN(2) == 0 {
		cam = camera.NewExternalLightCamera(id, an)
	} else {
		cam = camera.NewNightLightCamera(id, an)
	}
	return cam, an
}

func main() {
	cameraCount := rand.IntN(10) + 5

	var cameras []camera.Camera
	var animals []*animal.Animal

	for i := range cameraCount {
		cam, an := createAnimalAndCamera(i)
		cameras = append(cameras, cam)
		animals = append(animals, an)
	}

	fmt.Printf("Created %v animals and their cameras\n", len(animals))
	server := server.NewServer()

	movesCount := 10
	timeOfDay := camera.PartOfDay("Morning")
	for i := 0; i < movesCount; i++ {
		timeOfDay = camera.NextPartOfDay(timeOfDay)
		fmt.Printf("\nNext part of the day: %v, handling it\n", timeOfDay)

		fhEpisode := animal.CreateFullHistoryEpisode()
		dhEpisode := animal.CreateDangerousHistoryEpisode()
		for j := 0; j < cameraCount; j++ {
			animals[j].State = animal.GenerateRandomAnimalState()
			err := cameras[j].Process(timeOfDay)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// If animal state is not dangerous it won't be added to dhEpisode
			dhEpisode.Add(animals[j])
			fhEpisode.Add(animals[j])
		}
		server.FullHistory = append(server.FullHistory, fhEpisode)
		fmt.Printf("Dangerous history will be extended for: %v\n", dhEpisode.GetData())
		server.DangerousHistory = append(server.DangerousHistory, dhEpisode)
	}

	randEpisodeIndex := rand.IntN(movesCount)
	fmt.Printf("\nGet history for the move: %v\n", randEpisodeIndex)
	server.PrintCompleteHistoryForID(randEpisodeIndex)
}
