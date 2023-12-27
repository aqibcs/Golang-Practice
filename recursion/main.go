package main

import "fmt"

// Factorial calculates the factorial of a non-negative integer
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci calculates the nth Fibonacci number using recursion.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	// Example of factorial calculation
	fmt.Println("Factorial of 5:", Factorial(5))

	// Example of Fibonacci sequence
	fmt.Println("Fibonacci(8):", Fibonacci(8))
}
