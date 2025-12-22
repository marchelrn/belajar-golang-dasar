package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Marchel Manullang\n" +
		"Marchel,Matthew\n" +
		"Joko, Widodo"
	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
