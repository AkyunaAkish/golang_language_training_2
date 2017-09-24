package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePointer, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// io copy utilizes the reader interface
	// read method to turn a slice of bytes
	// into a readable format and
	// give it to a specified output
	// i.e. the standard out
	// of the current operating sytem
	io.Copy(os.Stdout, filePointer)
}

// Take file name as terminal arugment
// using os.Args
// then Read the contents of the file
// and print it to the command line
