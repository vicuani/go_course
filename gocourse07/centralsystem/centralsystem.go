package centralsystem

import (
	"context"
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

func (cs *CentralSystem) ProcessData(ctx context.Context, sensorChan <-chan sensor.SensorData, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received done signal, exiting from central system...")
			return
		case data, ok := <-sensorChan:
			if !ok {
				return
			}
			fmt.Printf("Central system: processed such a data: %v\n", data)
			cs.WriteToDB(data)
		default:
			time.Sleep(time.Duration(time.Millisecond * time.Duration(10)))
		}
	}
}

func (cs *CentralSystem) WriteToDB(data sensor.SensorData) {
	fmt.Printf("Starting writing data: %v to db...\n", data)
	time.Sleep(100 * time.Millisecond)
	cs.data = append(cs.data, data)
	fmt.Printf("Finished writing data: %v to db. Total data amount: %v\n", data, cs.DataSize())
}
