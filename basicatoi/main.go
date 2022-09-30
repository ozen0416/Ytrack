package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {

	var tmp int
	for i, _ := range s {
		tmp = tmp*10 + int(s[i]-'0')

	}
	return tmp
}
