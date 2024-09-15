package main

import (
	"fmt"
	"sync"
	"time"
)

type logger struct{}

func (l *logger) printLogs(lChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case lg, ok := <-lChan:
			if !ok {
				return
			}
			fmt.Printf("[LOG] %v\n", lg)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
