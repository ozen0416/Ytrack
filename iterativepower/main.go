package main

import "fmt"

func IterativePower(nb int, power int) int {

	result := 1

	if power < 0 {
		return 0
	}

	if power <= 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		result = nb * result

	}
	return result
}

func main() {
	result2 := IterativePower(2, 4)
	fmt.Printf("the result of the iterative function is: %v", result2)
}
