package main

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"os"
	"sync"
	"testing"
)

func TestFoundGPRSSignal(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	collar := NewCollar(rand.IntN(50)+30, float64(rand.IntN(10)+32), logger)

	collar.FoundGPRSSignal()
	if !collar.HasGPRSSignal() {
		t.Errorf("Expected GPRS signal to be true, got false")
	}
}

func TestLostGPRSSignal(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	collar := NewCollar(rand.IntN(50)+30, float64(rand.IntN(10)+32), logger)

	collar.LostGPRSSignal()
	if collar.HasGPRSSignal() {
		t.Errorf("Expected GPRS signal to be false, got true")
	}
}

func TestCollectSensorData(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	collar := NewCollar(rand.IntN(50)+30, float64(rand.IntN(10)+32), logger)

	const collectDataCount = 7

	dataChan := make(chan AnimalData, collectDataCount)
	var wg sync.WaitGroup

	wg.Add(1)
	go collar.CollectSensorData(collectDataCount, dataChan, &wg)

	wg.Wait()

	count := 0
	for range dataChan {
		count++
	}

	if count != collectDataCount {
		t.Errorf("Expected 10 data points, got %d", count)
	}
}

func TestTransmitData(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	collar := NewCollar(rand.IntN(50)+30, float64(rand.IntN(10)+32), logger)
	const collectDataCount = 5

	dataChan := make(chan AnimalData, collectDataCount*collectDataCount)
	var wg sync.WaitGroup

	t.Run("transmit", func(t *testing.T) {
		collar.FoundGPRSSignal()

		wg.Add(1)
		go collar.CollectSensorData(collectDataCount, dataChan, &wg)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go collar.TransmitData(ctx, dataChan)

		wg.Wait()
		if len(collar.storedData) != 0 {
			t.Errorf("Expected stored data to be empty after transmit, got %d items", len(collar.storedData))
		}
	})

	t.Run("store", func(t *testing.T) {
		collar.LostGPRSSignal()
		dataChan := make(chan AnimalData, collectDataCount)
		wg.Add(1)
		go collar.CollectSensorData(collectDataCount, dataChan, &wg)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go collar.TransmitData(ctx, dataChan)

		wg.Wait()
		if len(collar.storedData) == 0 {
			t.Errorf("Expected stored data to be non-empty, got %d items", len(collar.storedData))
		}
	})
}
