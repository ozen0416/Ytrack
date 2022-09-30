package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	runeArray := []rune(s)
	arrayCount := 0
	for i := range s {
		arrayCount = i + 1
	}
	return runeArray[arrayCount-1]

}
