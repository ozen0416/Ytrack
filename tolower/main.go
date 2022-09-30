package main

import "fmt"

func ToLower(s string) string {

	b := []rune(s)
	for a := 0; a < len(s); a++ {
		if s[a] >= 65 && s[a] <= 90 {
			b[a] = rune(s[a]) + 32
		}

	}
	return string(b)
}
func main() {
	fmt.Println(ToLower("HHHHHH"))
}
