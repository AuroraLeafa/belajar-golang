package main

import "fmt"

func main() {

	fmt.Println(len("Hello, World!"))

	name := "Testing"
	fmt.Println(name)

	var (
		firstName = "Refansyah"
		lastName  = "Sug"
	)
	fmt.Println(firstName, lastName)

	const (
		conFirstName = "Reff"
		conLastName  = "Sug"
	)
	fmt.Println(conFirstName, conLastName)

	type noKtp string
	var ktpUser noKtp = "12312312"
	var contoh string = "12345"
	var contohKTP = noKtp(contoh)
	fmt.Println("KTP User :", ktpUser)
	fmt.Println("ContohKTP :", contohKTP)
}
