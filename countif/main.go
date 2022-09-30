package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func CountIf(f func(string) bool, tab []string) int {

	var tmp int
	for _, i := range tab {
		if IsNumeric(i) {
			tmp += 1
		}
	}

	return tmp
}

func IsNumeric(s string) bool {

	for _, a := range s {
		if a < 48 || a > 57 {
			return false

		}

	}
	return true

}
