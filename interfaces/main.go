package main

import "fmt"

// the english and spanish bots both have a
// getGreeting function that they
// implement(because they match the receiver for their respective getGreeting function)
// which returns a string and therefore
// they impement the bot interface
// and are usable in the printGreeting function
// that requires an implementation of the bot interface
type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	jenkins := englishBot{}
	aztec := spanishBot{}

	printGreeting(jenkins)
	printGreeting(aztec)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// Interface Explanation:

// Interfaces allow for a function to be written
// and instead of writing multiple duplicate
// functions for different types
// you can write one function with
// an interface as the receiver or argument

// When creating a data type in your code,
// if your data type meets the criteria
// of an interface such as having a property name of type string
// for example

// That data type now implements that interface
// and therefore can have access to the functions
// that use that interface as a receiver or argument

////////////////////////////////////

// Concrete types
// can be applied to a new variable

// Interface type
// doesn't get set directly to a variable
// but a variable can "implement" an interface
// and therefore be compatible with functions
// that utilize the implemented interface
