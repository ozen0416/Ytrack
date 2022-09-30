package main

import "fmt"

func MakeRange(min, max int) []int {

	tmp := make([]int, max-min+1)
	for i := 0; i < max-min+1; i++ {
		tmp[i] = min + i
		if min >= max {
			return nil
		}
	}
	return tmp
}

func main() {
	size := 5
	total := MakeRange(size, 9)
	fmt.Println(total)

}
