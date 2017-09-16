package main

import "fmt"

func main() {
	rs := makeRangeSlice(1, 10)

	for _, n := range rs {
		if n%2 != 0 {
			fmt.Println(n, "odd")
		} else {
			fmt.Println(n, "even")
		}
	}
}

func makeRangeSlice(min, max int) []int {
	a := make([]int, max-min+1)

	for i := range a {
		a[i] = min + i
	}

	return a
}
