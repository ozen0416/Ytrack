package main

import "fmt"

func Join(strs []string, sep string) string {

	result := ""
	result += strs[0]
	for a := 1; a < len(strs); a++ {

		result += sep + strs[a]

	}
	return result

}

func main() {
	strs := []string{"Hello ", "how ", "are ", "you?"}
	fmt.Println(Join(strs, ":"))
}
