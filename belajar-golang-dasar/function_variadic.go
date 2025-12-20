package main

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(10, 10, 10, 10, 10, 10)
	println(result)
	numbers := []int{20, 20, 20, 20, 20}
	result = sumAll(numbers...)
	println(result)
}
