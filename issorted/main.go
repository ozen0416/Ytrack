package main

import "fmt"

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(edit, a1)
	result2 := IsSorted(edit, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, a []int) bool {

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}

	}
	return true

}

func edit(a, b int) int {

	for i := 0; i < a+b; i++ {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		if a == b {
			return 0
		}
	}
	return a
}
