package main

// package main means this package will create an executable file when running go build onm it

import "fmt"

func main() {
	// var card string = "Ace of spades"
	card := "Ace of spades" // := tells go to infer the datatype of the variable
	fmt.Println(card)
}
