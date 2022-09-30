package main

import "fmt"

func main() {
	s := "Hello world!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	var tmp string
	for _, i := range s {
		tmp = string(i) + tmp

	}
	return tmp
}
