package centralsystem

import (
	"sync"
	"testing"

	"github.com/vicuani/go_course/gocourse07/sensor"
)

func TestCentralSystem_ProcessData(t *testing.T) {
	cs := &CentralSystem{}
	centralChan := make(chan sensor.SensorData, 1)
	var wg sync.WaitGroup

	centralChan <- sensor.NewSensorData("temperature", 25)
	close(centralChan)

	wg.Add(1)
	go cs.ProcessData(centralChan, &wg)
	wg.Wait()

	if cs.DataSize() != 1 {
		t.Errorf("Expected data size (1) in central system, got %v", cs.DataSize())
	}
}
