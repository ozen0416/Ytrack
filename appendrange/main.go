package main

import "fmt"

func AppendRange(min, max int) []int {

	tmp := []int{}
	for i := min; i < max; i++ {
		tmp = append(tmp, i)
		if min >= max {
			return nil
		}
	}
	return tmp
}

func main() {
	size := 5
	total := AppendRange(size, 10)
	fmt.Println(total)

}
