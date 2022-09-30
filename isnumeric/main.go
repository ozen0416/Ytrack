package main

import "fmt"

func IsNumeric(s string) bool {

	for _, a := range s {
		if a < 48 || a > 57 {
			return false

		}

	}
	return true

}
func main() {
	fmt.Println(IsNumeric("01457896"))
}
