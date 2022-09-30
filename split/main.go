package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {

	var tmp []string
	var add int = 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) > len(s) {
			break
		}
		if sep == s[i:i+len(sep)] {
			tmp = append(tmp, s[add:i])
			add = i + len(sep)
		}

	}
	tmp = append(tmp, s[add:(len(s))])
	return tmp

}
