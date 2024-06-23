package main

import "fmt"

func getFullName() (string, string) {
	return "Refansyach", "Sugianto"
}

func main() {
	my, name := getFullName()
	fmt.Println(my, name))
}
