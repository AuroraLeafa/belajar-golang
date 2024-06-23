package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println("Background : ", background)
	fmt.Println("TODO : ", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println("Isi Context :", contextA)
	fmt.Println("Isi Context :", contextB)
	fmt.Println("Isi Context :", contextC)
	fmt.Println("Isi Context :", contextD)
	fmt.Println("Isi Context :", contextE)
	fmt.Println("Isi Context :", contextF)

	fmt.Println(contextC.Value("c"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	counter := 1
	go func() {
		defer close(destination)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("TIMEOUT!")
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	fmt.Println("waktu sekarang:", time.Now())
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("TIMEOUT!")
	fmt.Println("Waktu Sekarang:", time.Now())
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
