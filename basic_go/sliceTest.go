package main

import "fmt"

func main() {
	names := [...]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus"}
	slice := names[3:8]
	fmt.Println(slice)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

}
