package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println("Now without local : ", now)

	utc := time.Date(1993, time.February, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())
	fmt.Println("UTC Without Local", utc)
}
