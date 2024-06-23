package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Second * 5600
	duration2 := time.Millisecond * 10
	duration3 := time.Hour * 5
	durationRes := duration3 - duration1

	fmt.Println(durationRes)
	fmt.Printf("Duration : %d \n", duration2)
}
