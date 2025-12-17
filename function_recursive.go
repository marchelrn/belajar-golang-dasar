package main

// for loop function

func factorialInt(value int) int {

	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// function recursive

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	println(factorialInt(10))
	println(factorialRecursive(10))
}
