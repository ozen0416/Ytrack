package main

import "fmt"

func main() {
	result := IsPrime(4)
	fmt.Println(result)
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
