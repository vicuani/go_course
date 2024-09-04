package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Collar struct {
	data       *AnimalData
	gprsSignal bool
	storedData []*AnimalData
	mtx        sync.Mutex
}

func NewCollar() *Collar {
	return &Collar{
		data:       NewAnimalData(rand.IntN(50)+30, float64(rand.IntN(10)+32)),
		gprsSignal: false,
	}
}

func (c *Collar) CollectSensorData(dataChan chan<- AnimalData, wg *sync.WaitGroup) {
	defer wg.Done()

	breathingSensor := BreathingSensor{}
	soundSensor := SoundSensor{}

	for i := 0; i < 10; i++ {
		c.breathingData(breathingSensor)
		c.soundData(soundSensor)

		c.mtx.Lock()
		c.data.timestamp = time.Now()
		dataChan <- *c.data
		c.mtx.Unlock()

		time.Sleep(500 * time.Millisecond)
	}
	close(dataChan)
}

func (c *Collar) breathingData(sensor Sensor[int]) {
	bData := sensor.GenerateData()
	c.mtx.Lock()
	c.data.breathing = append(c.data.breathing, bData)
	c.mtx.Unlock()
}

func (c *Collar) soundData(sensor Sensor[int]) {
	sData := sensor.GenerateData()
	c.mtx.Lock()
	c.data.sounds = append(c.data.sounds, sData)
	c.mtx.Unlock()
}

func (c *Collar) TransmitData(ctx context.Context, dataChan <-chan AnimalData, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received done signal, exiting from transmit data...")
			return
		case data, ok := <-dataChan:
			if !ok {
				return
			}

			if c.gprsSignal {
				fmt.Printf("Transmit data: %v to server...\n", data)
				c.TransmitStoredData()
			} else {
				fmt.Printf("Store data: %v to internal memory...\n", data)
				c.mtx.Lock()
				c.storedData = append(c.storedData, &data)
				c.mtx.Unlock()
			}
		}
	}
}

func (c *Collar) TransmitStoredData() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, storedData := range c.storedData {
		fmt.Printf("Transmit stored data: %v to server...\n", storedData)
	}
	c.storedData = nil
}
