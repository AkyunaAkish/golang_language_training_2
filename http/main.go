package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Response: ", resp)
}

// the Body property of the response struct
// is an interface io.ReadCloser

// if an interface is specified as a property
// of a struct
// all that means is that particular property
// needs to be a type that implements that particular
// interface

// In the case of response.Body
// the body needs to implement both the Read and Close
// Interfaces that are the requirements of the io.ReadCloser interface
