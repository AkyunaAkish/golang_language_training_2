// check each website infinitely
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// make a channel which can have
	// strings passed through it
	c := make(chan string)

	for _, link := range links {
		// run checkSiteStatus
		// call in a new go routine
		// passing in the channel "c"
		// to ensure that we can read
		// the values produced in the
		// go routines
		// in the main function
		go checkSiteStatus(link, c)
	}

	// ranging over channel "c"
	// and receiving a channel value as "l"
	for l := range c {
		// checking statuses
		// every 5 seconds
		// passing l to the anonymous function
		// ensures that l will always
		// be the new string "l"
		// from the channel
		// and not refer to the previous
		// meaning of "l"
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkSiteStatus(l, c)
		}(l)
	}
}

func checkSiteStatus(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println("An issue occurred when calling: ", l)
		fmt.Println("Error: ", err)

		// send string onto passed in channel
		c <- l
		return
	}

	fmt.Println("Success", l)
	// send string onto passed in channel
	c <- l
}

// check statuses of each website
// once within a go routine

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	links := []string{
// 		"https://google.com",
// 		"https://facebook.com",
// 		"https://stackoverflow.com",
// 		"https://golang.org",
// 		"https://amazon.com",
// 	}

// 	// make a channel which can have
// 	// strings passed through it
// 	c := make(chan string)

// 	for _, link := range links {
// 		// run checkSiteStatus
// 		// call in a new go routine
// 		// passing in the channel "c"
// 		// to ensure that we can read
// 		// the values produced in the
// 		// go routines
// 		// in the main function
// 		go checkSiteStatus(link, c)
// 	}

// 	for i := 0; i < len(links); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func checkSiteStatus(l string, c chan string) {
// 	_, err := http.Get(l)

// 	if err != nil {
// 		fmt.Println("An issue occurred when calling: ", l)
// 		fmt.Println("Error: ", err)

// 		// send string onto passed in channel
// 		c <- "An issue occurred when calling: " + l
// 		return
// 	}

// 	// send string onto passed in channel
// 	c <- "Success when calling: " + l
// }
