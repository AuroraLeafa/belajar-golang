package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Refansyach Sugianto", "Ref"))
	fmt.Println(strings.Split("Refansyach Sugianto", " "))
	fmt.Println(strings.ToLower("Refansyach Sugianto"))
	fmt.Println(strings.ToUpper("Refansyach Sugianto"))
	fmt.Println(strings.Trim("     Refansyach   Sugianto   ", " "))
	fmt.Println(strings.ReplaceAll("Refan Refan Refan Refansyach Refansyach Sugianto", "Refan", "Muh"))

}
