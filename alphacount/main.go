package main

import "fmt"

func Alphacount(s string) int {

	count := 0
	for _, s := range s {
		if isAlphacount(s) {
			count++
		}
	}
	return count
}

func isAlphacount(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
func main() {
	s := "Hello 78 wold!  4455/"
	nb := Alphacount(s)
	fmt.Println(nb)
}
