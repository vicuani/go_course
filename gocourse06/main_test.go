package main

import (
	"io"
	"log/slog"
	"sync"
	"testing"
)

func TestEnclosureController(t *testing.T) {
	logger = slog.New(slog.NewTextHandler(io.Discard, nil))

	enclosureReqChan := make(chan EnclosureRequest)
	var wg sync.WaitGroup

	wg.Add(1)
	go enclosureController(enclosureReqChan, &wg)

	respChan := make(chan bool)

	for range maxEnclosuresAccess {
		enclosureReqChan <- EnclosureRequest{
			animalID:     1,
			isOpenAction: true,
			respChan:     respChan,
		}
		accessGranted := <-respChan
		if !accessGranted {
			t.Errorf("Expected access to enclosure")
		}
	}

	enclosureReqChan <- EnclosureRequest{
		animalID:     1,
		isOpenAction: true,
		respChan:     respChan,
	}
	accessGranted := <-respChan
	if accessGranted {
		t.Errorf("Expected access denied to enclosure")
	}

	close(enclosureReqChan)
	wg.Wait()
}
