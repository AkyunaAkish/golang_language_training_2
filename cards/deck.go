package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of deck
// which is a slice of strings
// deck will extend/inherit the properties of a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function print with a deck type
// as the receiver so that the deck type
// will "implement" this function and
// can have print called on any instance of the deck type
func (d deck) print() {
	// d is the instance of the deck type
	// that is implementing and invoking the print function
	// the d argument is similar to "this"
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns two values(will be returned as a slice)
// and the first return value takes from the beginning of the deck
// until by not including the index equal to the hand size
// the second return value takes from the index of the deck
// equal to the handSize until the very last item in the deck slice
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	// check for error
	if err != nil {
		fmt.Println("Error: ", err)
		// exit function with exit code 1 meaning something went wrong
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")
	// s(slice of strings) can be converted to a deck
	// because the underlying type of deck is a slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	// for the random function to be truly random
	// everytime the function runs, a new "seed" or "source"
	// for a data type rand needs to be generated in order to not
	// have predictably random output
	// so rand can have a new source generated and simply needs an int64
	// generated, in order to make a unique int64, we take the current
	// time and convert it to type int64 by calling UnixNano
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// get random Integer from 0 - the length of d - 1
		newPosition := r.Intn(len(d) - 1)

		// swap d[i] with d[newPosition]
		// and take what's at d[newPosition] and swap with the element at d[i]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
