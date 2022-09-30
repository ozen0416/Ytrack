package main

import "fmt"

func Fibonacci(index int) int {

	if index < 2 {

	}
	if index == 0 {
		return (0)

	}
	if index == 1 {
		return (1)
	}

	return Fibonacci(index-2) + Fibonacci(index-1)
}

func main() {
	result := Fibonacci(6)
	fmt.Printf("the result of the fibonacci sequence is: %v", result)
}
