package main

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Marchel"
	middleName = "Matthew"
	lastName = "Manullang"
	return firstName, middleName, lastName
}
func main() {
	firstName, middleName, lastName := getCompleteName()

	println(firstName, middleName, lastName)
}
