package main

import "fmt"

func main() {
	// 3 ways to declare a map with a key of type string and
	// value of type string

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b4f745", // always needs a trailing comma
	}

	// square braces are the only way to
	// target and set values in a map
	// dot syntax is only available in structs

	// set a value
	colors["white"] = "#ffffff"

	//edit a value
	// colors["white"] = "#ffffff22222"

	// delete a value
	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

// WHY USE A MAP INSTEAD OF A STRUCT

// map requires all keys to be same type and all values to be same type
// a struct can be more mixed with it's key/value pair types

// keys are indexed in a map which allows for iteration
// which is not available in a struct

// a map is a reference type
// a struct is a value type
// which means a struct requires more thought
// around passing pointers around

// Use a map when creating a collection
// of very closely related values
// Also use a map if you don't know
// the keys ahead of time, you can add and remove
// them freely unlike a struct which
// needs all keys pre-defined

// Structs are used a lot in Go,
// when you want to create a key/value
// mapping with rigid key names
