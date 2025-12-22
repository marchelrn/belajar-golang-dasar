package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Marchel", "Manullang"})
	_ = writer.Write([]string{"Joko", "Widodo"})
	writer.Flush()
}
