package main

import "fmt"

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

func DivMod(a int, b int, div *int, mod *int) {
	*mod = a % b
	*div = a / b
}
