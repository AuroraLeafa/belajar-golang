package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Time has Passed:", i)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	runtime.GOMAXPROCS(20)                // Untuk merubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1) //Thread default adalah -1
	fmt.Println("Total Thread: ", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines: ", totalGoroutines)

	group.Wait()

}
