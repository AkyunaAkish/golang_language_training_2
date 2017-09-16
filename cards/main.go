package main

// package main means this package will create an executable file when running go build onm it
func main() {
	// using the newDeck func from the deck.go file
	// we receive a new value of custom type deck
	// with some cards in it
	// then we use the custom print method that is implemented
	// by the custom deck tyoe(because the print method has deck as a receiver)
	// and we see all the contents of the deck
	cards := newDeck()
	// cards.print()

	// because deal returns two values
	// we can set variables to those value in order
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// saves cards as a list of strings separated by a comma
	// using the custom saveToFile deck method from the deck.go file
	cards.saveToFile("my_cards")

	// can be used to read from a file and convert file into type deck
	cardsFromFile := newDeckFromFile("my_cards")

	cardsFromFile.shuffle()
	cardsFromFile.print()
}

// TO RUN CODE YOU NEED TO COMBILE MAIN AND DECK GO FILES
// $ go run cards/main.go cards/deck.go
