package main

func main() {

	counter := 0

	increment := func() {
		println("Saya dipanggil")
		counter++
	}

	increment()
	increment()
	increment()

	println("Counter:", counter)
}
