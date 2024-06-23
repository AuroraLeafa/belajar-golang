package main

import (
	"fmt"
	"go-dasar/database"
)

func main() {
	db := database.GetDatabase()
	fmt.Println(db)
}
