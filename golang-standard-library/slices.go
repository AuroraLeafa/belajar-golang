package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"john", "Paul", "George", "Aries"}
	Ages := []int{35, 32, 24, 23, 12}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(Ages))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(Ages))
	fmt.Println(slices.Contains(names, "reff"))
	fmt.Println(slices.Index(names, "reff"))
	fmt.Println(slices.Index(names, "Paul"))

}
