package main

import "fmt"

func ToUpper(s string) string {

	result := ""
	for _, a := range s {
		tmp := a
		if a >= 97 && a <= 122 {
			tmp = a - 32
		}
		result += string(tmp)
	}
	return result
}
func main() {
	fmt.Println(ToUpper("HytoH"))
}
