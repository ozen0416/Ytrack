package main

import "fmt"

func IterativeFactorial(index int) int {

	if index < 0 || index > 20 {
		return 0
	}

	result := 1

	for i := 1; i < index+1; i++ {
		result = result * i

	}
	return result
}

func main() {
	result2 := IterativeFactorial(4)
	fmt.Printf("the result of the iterative function is: %v", result2)
}
