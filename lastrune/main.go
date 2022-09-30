package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!  "))
	z01.PrintRune('\n')
}
func FirstRune(s string) rune {
	runeArray := []rune(s)
	arrayCount := 0
	for i := range s {
		arrayCount = i - 1
	}
	return runeArray[arrayCount-4]

}
