package main

func main() {
	runApplication()
}

func logging() {
	println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	println("Run Application")
}
