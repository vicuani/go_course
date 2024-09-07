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

type Camera interface {
	Process(pod camera.PartOfDay) error
}

func main() {
	var cameras []Camera
	var animals []*animal.Animal

	for i := range rand.IntN(10) + 5 {
		var cam Camera
		an := animal.NewAnimal(i)
		if rand.IntN(2) == 0 {
			cam = camera.NewExternalLightCamera(i, an)
		} else {
			cam = camera.NewNightLightCamera(i, an)
		}

		cameras = append(cameras, cam)
		animals = append(animals, an)
	}

	fmt.Printf("Created %v animals and their cameras\n", len(animals))
	srv := server.NewServer()

	movesCount := 10
	timeOfDay := camera.PartOfDay("Morning")
	for i := 0; i < movesCount; i++ {
		timeOfDay, err := camera.NextPartOfDay(timeOfDay)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nNext part of the day: %v, handling it\n", timeOfDay)

		fhEpisode := server.CreateFullHistoryEpisode()
		dhEpisode := server.CreateDangerousHistoryEpisode()
		for j := 0; j < len(cameras); j++ {
			animals[j].SetRandomState()
			err := cameras[j].Process(timeOfDay)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// If animal state is not dangerous it won't be added to dhEpisode
			dhEpisode.Add(animals[j])
			fhEpisode.Add(animals[j])
		}
		srv.AddFullHistoryEpisode(fhEpisode)
		fmt.Printf("Dangerous history will be extended for: %v\n", dhEpisode.GetData())
		srv.AddDangerousHistoryEpisode(dhEpisode)
	}

	randEpisodeIndex := rand.IntN(movesCount)
	fmt.Printf("\nGet history for the move: %v\n", randEpisodeIndex)
	srv.PrintCompleteHistoryForID(randEpisodeIndex)
}
