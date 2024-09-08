package sensor

import (
	"sync"
	"testing"
)

func TestBrightnessSensor_CollectData(t *testing.T) {
	centralChan := make(chan SensorData, 1)
	var wg sync.WaitGroup
	sensor := NewBrightnessSensor(0, 100)

	wg.Add(1)
	go sensor.CollectData(centralChan, 1, &wg)

	wg.Wait()
	close(centralChan)

	if len(centralChan) != 1 {
		t.Errorf("Expected 1 piece of data, got %v", len(centralChan))
	}

	data := <-centralChan
	if data.sensorType != "brightness" {
		t.Errorf("Expected sensor type 'brightness', got %v", data.sensorType)
	}
}
