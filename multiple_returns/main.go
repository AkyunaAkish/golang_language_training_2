package main

import "fmt"

func main() {
	c1, c2, c3 := colors()
	fmt.Println(c1, c2, c3)
}

func colors() (string, string, string) {
	return "red", "yellow", "blue"
}
