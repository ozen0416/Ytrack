package main

import "fmt"

func IsNumeric(s string) bool {

	for _, a := range s {
		if a < 48 || a > 57 && a < 65 || a > 90 && a < 97 || a > 122 {
			return false

		}

	}
	return true

}
func main() {
	fmt.Println(IsNumeric("whatsup89"))
}
