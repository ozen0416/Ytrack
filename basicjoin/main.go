package main

import "fmt"

func BasicJoin(elems []string) string {

	result := ""
	for _, a := range elems {

		result += a

	}
	return result

}

func main() {
	elems := []string{"Hello ", "how ", "are ", "you?"}
	fmt.Println(BasicJoin(elems))
}
