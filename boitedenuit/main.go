package main

import "fmt"

func CheckAge(nb int) bool {

	if nb >= 18 {
		return true
	}
	return false

}

func CheckIdentity(b string, age int) {

	if CheckAge(age) && b == "femme" {
		fmt.Print("Tu peux entrer, c'est 10 €")
	} else if CheckAge(age) && b == "homme" {
		fmt.Print("Tu peux entrer, c'est 15 €")

	} else {
		fmt.Print("Sortez!")
	}

}

func main() {
	CheckIdentity("femme", 7)
}
