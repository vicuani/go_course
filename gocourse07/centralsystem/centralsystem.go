package centralsystem

import (
	"fmt"
	"sync"
	"time"

	"github.com/vicuani/go_course/gocourse07/sensor"
)

type CentralSystem struct {
	data []sensor.SensorData
}

func (cs *CentralSystem) DataSize() int {
	return len(cs.data)
}

func (cs *CentralSystem) ProcessData(sensorChan <-chan sensor.SensorData, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-sensorChan:
			if !ok {
				return
			}
			fmt.Printf("Central system: processed such a data: %v\n", data)
			cs.writeToDB(data)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (cs *CentralSystem) writeToDB(data sensor.SensorData) {
	fmt.Printf("Starting writing data: %v to db...\n", data)
	time.Sleep(100 * time.Millisecond)
	cs.data = append(cs.data, data)
	fmt.Printf("Finished writing data: %v to db. Total data amount: %v\n", data, cs.DataSize())
}
