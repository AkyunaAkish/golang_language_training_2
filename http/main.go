package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Example of letting Golang
	// convert Body to byte slice and convert
	// to string and then printing it to
	// the terminal
	// the io.Copy method implements
	// the Writer interface
	// and therefore can convert
	// data from a Reader interface
	// and write it to an output

	// custom writer using log writer
	io.Copy(lw, resp.Body)

	// using the operating system's
	// standard out as the data type
	// that implements the Writer interface
	// (which will print to the terminal)
	// io.Copy(os.Stdout, resp.Body)

	// Example of manually reading from the
	// resp body and printing it out as
	// a string:

	// // make takes a type of a slice
	// // and then allows you to initialize
	// // an amount of elements to start out with
	// byteSlice := make([]byte, 99999)

	// // the resp body needs to have it's data
	// // added to a slice of bytes using the read
	// // method before you can convert it to a
	// // string and Println it
	// resp.Body.Read(byteSlice)

	// fmt.Println("byteSlice: ", string(byteSlice))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes: ", len(bs))
	return len(bs), nil
}

// the Body property of the response struct
// is an interface io.ReadCloser

// if an interface is specified as a property
// of a struct
// all that means is that particular property
// needs to be a type that implements that particular
// interface

// In the case of response.Body
// the body needs to implement both the Reader and Close
// Interfaces that are the requirements of the io.ReadCloser interface

// Reader Interface
// Every data type implements the ready interface
// and therefore is compatible with the res.Body property
// the Read function is applied to res.Body which
// is why all data coming through the http
// request can be used as res.Body
