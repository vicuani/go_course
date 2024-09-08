package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestSetGPRSSignal(t *testing.T) {
	collar := NewCollar()

	t.Run("testFalse", func(t *testing.T) {
		collar.SetGPRSSignal(false)
		if collar.GPRSSignal() {
			t.Errorf("Expected GPRS signal to be false, got true")
		}
	})

	t.Run("testTrue", func(t *testing.T) {
		collar.SetGPRSSignal(true)
		if !collar.GPRSSignal() {
			t.Errorf("Expected GPRS signal to be true, got false")
		}
	})
}

func TestCollectSensorData(t *testing.T) {
	collar := NewCollar()
	dataChan := make(chan AnimalData, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go collar.CollectSensorData(dataChan, &wg)

	wg.Wait()

	count := 0
	for range dataChan {
		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 data points, got %d", count)
	}
}

func TestTransmitData(t *testing.T) {
	collar := NewCollar()
	collar.SetGPRSSignal(true)

	dataChan := make(chan AnimalData, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go collar.CollectSensorData(dataChan, &wg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go collar.TransmitData(ctx, dataChan)

	time.Sleep(1 * time.Second)

	collar.SetGPRSSignal(true)
	wg.Wait()
}
