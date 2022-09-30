package main

import (
	"fmt"
	"strings"
)

func Index(s string, toFind string) int {

	result1 := strings.Index("hello", "4")
	return result1

}

func main() {
	result1 := Index("Hello", "o")
	fmt.Println(result1)

}
