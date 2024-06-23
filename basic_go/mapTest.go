package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
		"NoHP":    "123456",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["NoHP"])
	fmt.Println(len(person))
	delete(person, "NoHP")
	fmt.Println(person["NoHP"])
}
