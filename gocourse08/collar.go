package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Collar struct {
	data *AnimalData

	gprsSignalMu sync.Mutex
	gprsSignal   bool

	storedDataMu sync.Mutex
	storedData   []*AnimalData
}

func NewCollar() *Collar {
	return &Collar{
		data:       NewAnimalData(rand.IntN(50)+30, float64(rand.IntN(10)+32)),
		gprsSignal: false,
	}
}

func (c *Collar) SetGPRSSignal(v bool) {
	c.gprsSignalMu.Lock()
	c.gprsSignal = v
	c.gprsSignalMu.Unlock()
}

func (c *Collar) GPRSSignal() bool {
	defer c.gprsSignalMu.Unlock()

	c.gprsSignalMu.Lock()
	return c.gprsSignal
}

func (c *Collar) CollectSensorData(dataChan chan<- AnimalData, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10 {
		c.breathingData(BreathingSensor{})
		c.soundData(SoundSensor{})

		c.data.timestamp = time.Now()
		dataChan <- *c.data

		time.Sleep(500 * time.Millisecond)
	}
	close(dataChan)
}

func (c *Collar) breathingData(sensor Sensor[int]) {
	c.data.breaths = append(c.data.breaths, sensor.GenerateData())
}

func (c *Collar) soundData(sensor Sensor[int]) {
	c.data.sounds = append(c.data.sounds, sensor.GenerateData())
}

func (c *Collar) TransmitData(ctx context.Context, dataChan <-chan AnimalData) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received done signal, exiting from transmit data...")
			return
		case data, ok := <-dataChan:
			if !ok {
				return
			}

			if c.GPRSSignal() {
				fmt.Printf("Transmit data: %v to server...\n", data)
				c.TransmitStoredData()
			} else {
				fmt.Printf("Store data: %v to internal memory...\n", data)
				c.storedDataMu.Lock()
				c.storedData = append(c.storedData, &data)
				c.storedDataMu.Unlock()
			}
		}
	}
}

func (c *Collar) TransmitStoredData() {
	defer c.storedDataMu.Unlock()

	c.storedDataMu.Lock()
	for _, storedData := range c.storedData {
		fmt.Printf("Transmit stored data: %v to server...\n", storedData)
	}
	c.storedData = nil
}
