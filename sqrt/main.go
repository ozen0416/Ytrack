package main

import "fmt"

func Sqrt(nb int) int {

	for a := 1; nb/a >= a; a++ {
		if nb/a == a {
			return a
		}

	}

	return 0

}
func main() {
	fmt.Println(Sqrt(81))
}
