package main

import "fmt"

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	var IsNegative bool
	if s[0] == 45 {
		IsNegative = true
	}
	var start int
	if s[0] == 45 || s[0] == 43 {
		start++
	}

	var tmp int
	for i := start; i < len(s); i++ {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return 0
		}
		tmp = tmp*10 + int(rune(s[i])-48)

	}
	if IsNegative {
		tmp = -tmp
	}
	return tmp

}
