package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n > 9 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + 48))

}
