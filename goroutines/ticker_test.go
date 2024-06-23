package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		ticker.Stop()
	}()
	for time := range ticker.C {
		fmt.Println("Ticker Now: ", time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println("Ticker Now: ", time)
	}
}
