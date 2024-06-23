package main

import "fmt"

func main() {
	type Customer struct {
		a
		Name, Address string
		Age           int
	}

	eko := Customer{"Eko", "Bogor", 24}
	fmt.Println(eko)
}
