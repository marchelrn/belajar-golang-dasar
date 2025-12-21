package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.December, 17, 12, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	dataList := []string{

		"2025-02-12 13:00:10",
		"2023-05-01 09:30:00",
		"2024-12-25 18:45:30",
	}

	data := dataList[1]
	dataTime, err := time.Parse(formatter, data)
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(dataTime)
	}
	fmt.Println(dataTime.Date())
	fmt.Println(dataTime.Hour())
	fmt.Println(dataTime.Minute())
	fmt.Println(dataTime.Second())
	fmt.Println(dataTime.Nanosecond())
}
