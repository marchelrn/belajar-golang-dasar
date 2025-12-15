package main

import "fmt"

func main() {
	userScore := 79
	userPresence := 14
	userFirstname := "Marchel"
	userLastname := "Manullang"

	if eligible(userScore, userPresence) == true {
		fmt.Println(graduate(userFirstname, userLastname))
	} else {
		fmt.Println("see you next year")
	}
}

func eligible(scores int, presence int) bool {
	if scores >= 80 && presence >= 13 {
		return true
	}
	return false
}

func graduate(firstName, lastName string) string {
	grats := "Congratulations " + firstName + " " + lastName

	return grats
}
