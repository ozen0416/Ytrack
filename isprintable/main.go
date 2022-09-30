package main

import "fmt"

func IsPrintable(s string) bool {

	for _, a := range s {
		if a < 33 || a > 126 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(IsPrintable("Hello\n"))
}
