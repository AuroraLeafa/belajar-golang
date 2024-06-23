package main

import "fmt"

func main() {
	name := "Reff"
	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Reff" {
		fmt.Println("Hello Reff")
	} else if name == "A" {
		fmt.Println("Hello A")
	} else if name == "B" {
		fmt.Println("Hello B")
	} else {
		fmt.Println("Boleh Kenalan ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}

}
