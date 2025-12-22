package main

import (
	"bufio"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("THIS IS LONG STRING \nwith new lines \n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		println(string(line))
	}
}
