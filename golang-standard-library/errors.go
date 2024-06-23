package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	ExceptionError  = errors.New("Exception error")
	NotFoundError   = errors.New("Not Found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "reff" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("budi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else if errors.Is(err, ExceptionError) {
			fmt.Println("Exception Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
