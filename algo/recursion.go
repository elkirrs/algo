package algo

import "fmt"

// Factorial
func Factorial(n int) int {
	iteration := 0
	result := factorialRun(n, &iteration)
	fmt.Printf("factorial: value: %d iterations: %d\n", result, iteration)
	return result
}

func factorialRun(n int, iterations *int) int {
	*iterations++
	if n == 1 {
		return 1
	}
	return n * factorialRun(n-1, iterations)
}

// Fibonachi
func Fibonachi(n int) int {
	iteration := 0
	result := fibonachiRun(n, &iteration)
	fmt.Printf("fibonachi: value: %d iterations: %d\n", result, iteration)
	return result
}

func fibonachiRun(n int, iterations *int) int {
	*iterations++
	if n == 1 || n == 2 {
		return 1
	}
	return fibonachiRun(n-1, iterations) + fibonachiRun(n-2, iterations)
}
