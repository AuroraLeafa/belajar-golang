package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Muhammad", "Refansyach", "Sugianto"})
	_ = writer.Write([]string{"Sugianto", "Muhammad", "Refansyach"})
	_ = writer.Write([]string{"Refansyach", "Sugianto", "Muhammad"})
	writer.Flush()
}
