package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {

	var tmp int
	for i, _ := range s {
		if s[i] == 32 {
			return 0
		}
		if s[i] >= 65 && s[i] <= 90 || s[i] >= 97 && s[i] <= 122 {
			return 0
		}
		tmp = tmp*10 + int(s[i]-'0')

	}
	return tmp
}
