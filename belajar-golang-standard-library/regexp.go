package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`m([a-z])+el`)

	fmt.Println(regex.MatchString("Marchel"))
	fmt.Println(regex.FindAllString("march marchel marchek mini munu meno minel minsel marcel michel michelle mansel msel mwensel mapel manel", 11))
}
