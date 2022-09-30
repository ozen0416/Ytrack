package main

import "fmt"

func Concat(str1 string, str2 string) string {

	str1 = "Hello!"
	str2 = "How are you ?"

	result := str1 + str2

	return result
}
func main() {
	result1 := Concat("Hello!", "How are you ?")
	fmt.Println(result1)
}
