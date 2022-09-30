package main

import "fmt"

func ConcatParams(args []string) string {

	tmp := ""

	for i := 0; i < len(args); i++ {
		tmp += "\n" + args[i]
	}
	return tmp
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
