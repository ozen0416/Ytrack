package main

import "fmt"

func IsUpper(s string) bool {

	for _, a := range s {
		if a < 97 || a > 122 {
			return false

		}

	}
	return true

}
func main() {
	fmt.Println(IsUpper("yes"))
}
