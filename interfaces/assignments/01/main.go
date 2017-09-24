package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   3,
		height: 5,
	}

	s := square{
		side: 8,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
	// return (t.base * t.height) * 0.5 // alternative equation
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
