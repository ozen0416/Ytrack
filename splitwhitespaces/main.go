package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))

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
