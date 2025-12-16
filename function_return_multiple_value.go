package main

func getFullname() (string, string) {
	return "Marchel", "Manullang"
}

func main() {
	firstName, lastName := getFullname()

	println(firstName, lastName)
}
