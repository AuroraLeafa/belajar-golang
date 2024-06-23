package main

import "fmt"

func Ups() any {
	return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
