/*

Програма "Вольєр для екзотичних птахів"

Вам потрібно створити програму для управління вольєром для екзотичних птахів у розумному зоопарку.
Вольєр має сенсори, які вимірюють температуру, яскравість освітлення та вологість,
і вони передають ці дані до центральної системи, яка, в свою чергу, зберігає їх у памʼять (така от база даних).

Раз на добу сенсори відключаються для технічної перевірки, але потім знову продовжують працювати.
Центральна система також перезавантажується раз на добу.

Робота кожного сенсора — окрема горутина.
Коли сенсори відключаються, горутини безпечно вимикається.
Те саме з центральною системою.
Процес запису в памʼять повинен бути (штучно) тривалим.
І у випадку, коли центральна система планово відключається,
вона це мусить робити лише після того, як всі записи в базу виконались.

*/

package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/vicuani/go_course/gocourse07/centralsystem"
	"github.com/vicuani/go_course/gocourse07/sensor"
)

const daysCount = 4
const maxIterations = 10

func main() {

	sensors := []sensor.Sensor{
		sensor.CreateBrightnessSensor(0, 100),
		sensor.CreateHumiditySensor(0, 100),
		sensor.CreateTemperatureSensor(10, 40),
	}

	//	make a channel with buffer able to contain all data from all sensors during day
	centralSystem := centralsystem.CentralSystem{}

	var totalIterationsCount int
	for i := range daysCount {
		centralChan := make(chan sensor.SensorData, len(sensors)*maxIterations)

		var sensorsWg sync.WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		fmt.Printf("\n*** Starting day #%v ****\n", i+1)
		var dailyIterationsCount int

		for _, s := range sensors {
			sensorsWg.Add(1)
			iterationsCount := rand.IntN(maxIterations-5) + 5
			dailyIterationsCount += iterationsCount
			go s.CollectData(centralChan, iterationsCount, &sensorsWg)
		}

		var centralWg sync.WaitGroup
		centralWg.Add(1)
		go centralSystem.ProcessData(ctx, centralChan, &centralWg)

		sensorsWg.Wait()
		close(centralChan)
		centralWg.Wait()

		fmt.Printf("During this day all sensors generated %v data, total consumed: %v\n", dailyIterationsCount, centralSystem.DataSize())
		totalIterationsCount += dailyIterationsCount
		fmt.Printf("\n*** Day #%v finished ****\n", i+1)
	}

	fmt.Printf("All processing is finished: generated data count = %v, total consumed data = %v\n", totalIterationsCount, centralSystem.DataSize())
}
