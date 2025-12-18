package main

func main() {
	runApp(true)
}

func endApp() {
	println("endApp")
	message := recover()
	println("Error : ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
}
