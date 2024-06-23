package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Muhammad, Refansyach, Sugianto\n" +
		"Sugianto, Muhammad, Refansyach\n" +
		"Refansyach, Sugianto, Muhammad"

	reader := csv.NewReader(strings.NewReader((csvString)))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
