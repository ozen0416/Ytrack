package main

import "fmt"

func RecursivePower(nb int, power int) int {

	if power < 0 {
		return 0

	}

	if power == 0 {
		return 1
	}

	return nb * RecursivePower(nb, power-1)
}

func main() {
	result := RecursivePower(2, 3)
	fmt.Printf("the result of the recursive function is: %v", result)
}
