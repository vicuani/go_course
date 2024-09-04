package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func SimulateGPRSSignal(collar *Collar, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Set GPRS signal to false")
		collar.gprsSignal = false

		absenseTime := rand.IntN(100) + 50
		time.Sleep(time.Millisecond * time.Duration(absenseTime))

		fmt.Println("Set GPRS signal to true")
		collar.gprsSignal = true

		presenseTime := rand.IntN(1000) + 500
		time.Sleep(time.Millisecond * time.Duration(presenseTime))
	}
}

func main() {
	collar := NewCollar()

	var wg sync.WaitGroup

	wg.Add(1)
	go SimulateGPRSSignal(collar, &wg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataChan := make(chan AnimalData, 100)
	wg.Add(1)
	go collar.CollectSensorData(dataChan, &wg)

	wg.Add(1)
	go collar.TransmitData(ctx, dataChan, &wg)

	wg.Wait()
	fmt.Println("The end")
}
