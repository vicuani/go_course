package main

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"os"
	"sync"
	"time"
)

var logger *slog.Logger

const sensorDataCollectsCount = 10

func SimulateGPRSSignal(collar *Collar, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 5 {
		logger.Info("Lost GPRS signal")
		collar.LostGPRSSignal()

		absenceTime := rand.IntN(100) + 50
		time.Sleep(time.Millisecond * time.Duration(absenceTime))

		logger.Info("Found GPRS signal")
		collar.FoundGPRSSignal()

		presenseTime := rand.IntN(1000) + 500
		time.Sleep(time.Millisecond * time.Duration(presenseTime))
	}
}

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	collar := NewCollar(rand.IntN(50)+30, float64(rand.IntN(10)+32), logger)

	var wg sync.WaitGroup

	wg.Add(1)
	go SimulateGPRSSignal(collar, &wg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataChan := make(chan AnimalData, 100)
	wg.Add(1)
	go collar.CollectSensorData(sensorDataCollectsCount, dataChan, &wg)

	collar.TransmitData(ctx, dataChan)

	wg.Wait()
	logger.Info("The end")
}
