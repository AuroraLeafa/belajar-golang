package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var (
		a = 2
		b = 3
		c = 4
		d = 5
	)
	a += b
	b -= c
	c *= d
	d /= a
	fmt.Println(a, b, c, d)
}
