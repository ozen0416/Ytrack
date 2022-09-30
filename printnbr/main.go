package main

import "github.com/01-edu/z01"

func Nbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n > 9 {
		Nbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + 48))

}

func main() {
	Nbr(-8)
}
