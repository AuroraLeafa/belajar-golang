package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number Now : ", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 50000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
