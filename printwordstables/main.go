package main

import "github.com/01-edu/z01"

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

func PrintWordsTables(a []string) {
	for _, z := range a {
		for _, x := range z {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	}
}

func SplitWhiteSpaces(s string) []string {

	var tmp []string
	var add int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			tmp = append(tmp, s[add:i])
			add = i + 1
		}

	}
	tmp = append(tmp, s[add:(len(s))])
	return tmp
}
