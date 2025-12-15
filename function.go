package main

import "fmt"

func main() {
	user_score := 79
	user_presence := 14
	user_firstName := "Marchel"
	user_lastName := "Manullang"

	if eligible(user_score, user_presence) == true {
		fmt.Println(graduate(user_firstName, user_lastName))
	} else {
		fmt.Println("see you next year")
	}
}

func eligible(scores int, presence int) bool {
	if scores >= 80 && presence >= 13 {
		return true
	} else {
		return false
	}
}

func graduate(firstName, lastName string) string {
	grats := "Congratulations " + firstName + " " + lastName

	return grats
}
