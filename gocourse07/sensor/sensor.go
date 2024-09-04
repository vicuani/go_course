package sensor

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const minInterval = 10
const maxInterval = 50

type SensorData struct {
	Type  string
	Value int
}

func CreateSensorData(t string, v int) SensorData {
	return SensorData{
		Type:  t,
		Value: v,
	}
}

type Sensor interface {
	CollectData(centralCh chan<- SensorData, iterationsCount int, wg *sync.WaitGroup)
}

func collect(centralChan chan<- SensorData, name string, minValue int, maxValue int, iterationsCount int) {
	for range iterationsCount {
		data := rand.IntN(maxValue-minValue) + minValue
		fmt.Printf("\nCollect data for %v: %v\n", name, data)
		centralChan <- CreateSensorData(name, data)
		time.Sleep(time.Duration(rand.IntN(maxInterval-minInterval)+minInterval) * time.Millisecond)
	}
}

type BrightnessSensor struct {
	name     string
	minValue int
	maxValue int
}

func CreateBrightnessSensor(minValue int, maxValue int) Sensor {
	return &BrightnessSensor{
		name:     "brightness",
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (bs *BrightnessSensor) CollectData(centralChan chan<- SensorData, iterationsCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	collect(centralChan, bs.name, bs.minValue, bs.maxValue, iterationsCount)
}

type HumiditySensor struct {
	name     string
	minValue int
	maxValue int
}

func CreateHumiditySensor(minValue int, maxValue int) Sensor {
	return &HumiditySensor{
		name:     "humidity",
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (hs *HumiditySensor) CollectData(centralChan chan<- SensorData, iterationsCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	collect(centralChan, hs.name, hs.minValue, hs.maxValue, iterationsCount)
}

type TemperatureSensor struct {
	name     string
	minValue int
	maxValue int
}

func CreateTemperatureSensor(minValue int, maxValue int) Sensor {
	return &TemperatureSensor{
		name:     "temperature",
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (ts *TemperatureSensor) CollectData(centralChan chan<- SensorData, iterationsCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	collect(centralChan, ts.name, ts.minValue, ts.maxValue, iterationsCount)
}
