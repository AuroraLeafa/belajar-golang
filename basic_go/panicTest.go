package main

import "fmt"

func endApp() {
	fmt.Println("End App!")
	message := recover()
	fmt.Println("Error message :", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups... Error Occured!")
	}
}

func main() {
	runApp(true)
	fmt.Println("TEST")
}
