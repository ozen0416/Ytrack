package main

import "fmt"

func RecursiveFactorial(index int) int {

	if index < 0 || index > 20 {
		return 0
	}

	if index <= 1 {
		return 1
	}

	return index * RecursiveFactorial(index-1)

}

func main() {
	result := RecursiveFactorial(4)
	fmt.Printf("the result of the recursive function is: %v", result)
}
