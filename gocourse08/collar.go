package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type Collar struct {
	pulse       int
	temperature float64

	gprsSignalMu sync.Mutex
	gprsSignal   bool

	storedDataMu sync.Mutex
	storedData   []*AnimalData

	logger *slog.Logger
}

func NewCollar(pulse int, temperature float64, logger *slog.Logger) *Collar {
	return &Collar{
		pulse:       pulse,
		temperature: temperature,
		gprsSignal:  false,
		logger:      logger,
	}
}

func (c *Collar) FoundGPRSSignal() {
	c.gprsSignalMu.Lock()
	c.gprsSignal = true
	c.gprsSignalMu.Unlock()
}

func (c *Collar) LostGPRSSignal() {
	c.gprsSignalMu.Lock()
	c.gprsSignal = false
	c.gprsSignalMu.Unlock()
}

func (c *Collar) HasGPRSSignal() bool {
	defer c.gprsSignalMu.Unlock()

	c.gprsSignalMu.Lock()
	return c.gprsSignal
}

func (c *Collar) CollectSensorData(collectDataCount int, dataChan chan<- AnimalData, wg *sync.WaitGroup) {
	defer wg.Done()

	data := NewAnimalData(c.pulse, c.temperature)
	for range collectDataCount {
		data.breaths = append(data.breaths, BreathingSensor{}.Measure())
		data.sounds = append(data.sounds, SoundSensor{}.Measure())

		data.timestamp = time.Now()
		dataChan <- *data

		time.Sleep(500 * time.Millisecond)
	}
	close(dataChan)
}

func (c *Collar) TransmitData(ctx context.Context, dataChan <-chan AnimalData) {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Received done signal, exiting from transmit data...")
			return
		case data, ok := <-dataChan:
			if !ok {
				return
			}

			if c.HasGPRSSignal() {
				c.logger.Info(fmt.Sprintf("Transmit data: %v to server...", data))
				c.TransmitStoredData()
			} else {
				c.logger.Info(fmt.Sprintf("Store data: %v to internal memory...", data))
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
		c.logger.Info(fmt.Sprintf("Transmit stored data: %v to server...", storedData))
	}
	c.storedData = nil
}
