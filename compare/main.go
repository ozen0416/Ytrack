package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) int {

	a = "H"
	b = "H"

	v := strings.Compare(a, b)

	return v
}

func main() {
	result := Compare("Hello", "Hi")
	fmt.Println(result)
}
