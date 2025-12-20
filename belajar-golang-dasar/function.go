package main

import "fmt"

func main() {
	userScore := 80
	userPresence := 13
	userFirstname := "Marchel"
	userLastname := "Manullang"

	sayHello()

	if eligible(userScore, userPresence) == true {
		fmt.Println(graduate(userFirstname, userLastname))
		fmt.Println("Your score :", grade(userScore))
	} else {
		fmt.Println("see you next year")
		fmt.Println("Your Score :", grade(userScore))
	}
}

// Function
func sayHello() {
	fmt.Println("Hello")
}

// Function with parameter and return value
func grade(scores int) string {
	if scores >= 80 {
		return "A"
	} else if scores >= 75 {
		return "B+"
	} else if scores >= 70 {
		return "B"
	} else if scores >= 65 {
		return "C+"
	} else if scores >= 60 {
		return "C"
	} else if scores >= 50 {
		return "D"
	}
	return "E"
}

func eligible(scores int, presence int) bool {
	if scores >= 60 && presence >= 13 {
		return true
	}
	return false
}

func graduate(firstName, lastName string) string {
	grats := "Congratulations " + firstName + " " + lastName

	return grats
}
