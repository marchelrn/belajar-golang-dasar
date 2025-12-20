package main

func sayGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	gb := sayGoodBye

	println(gb("Marchel"))
}
