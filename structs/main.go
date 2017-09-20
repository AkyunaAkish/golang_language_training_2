package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// alternative way to declare
// contactInfo(will automatically set field "contactInfo" of type contactInfio
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

func main() {
	// alternative way to declare a person
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "alex.anderson@gmail.com"
	// alex.contact.zipCode = 80303

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 80303,
		},
	}

	// using the & will create a reference
	// to the main value pointer
	// so that the function "updateName"
	// will be able to mutate the original
	// value and not a copy of it
	// alexPointer := &alex
	// alexPointer.updateName("Alexander")
	// alex.print()

	// you don't actually need to create the reference
	// pointer using the & syntax
	// using the *person in the updateName
	// function is enough for Go to assume you mean
	// to reference the memory pointer of the person in context
	alex.updateName("Alexander")
	alex.print()

	// example of printing specific fields
	// fmt.Println(alex.firstName, alex.lastName)

	// a person can also be declared with the field names based on
	// order of original struct declaration(not recommended)
	// alex := person{"Alex", "Anderson"}
	// fmt.Println(alex.firstName, alex.lastName)
}

func (p person) print() {
	// this will print the entire p struct instance
	// injecting the full value into the %+v's place
	fmt.Printf("%+v", p)
}

// Golang is a "Pass by value" language
// which means by default, a copy of the original
// argument or struct receiver will be passed into
// functions
// you can explicitly tell to go pass the original pointer
// to the original passed in value to your function
// in order to allow for mutation of the passed arugment or receiver

// using the * next to the type of argument/receier means
// that "p" will show up in the function as the memory pointer's
// value. If the * wasn't there, because we're using a memory address
// pointer above using the & syntax, we would see the raw memory address

func (p *person) updateName(newFirstName string) {
	// fmt.Println(&p) // converting memory pointer's value back to raw memory address

	// mutating the pointer to person
	p.firstName = newFirstName

	// alternative syntax explicitly saying
	// we want to manipulate the value we're referencing
	// more directly, *P just makes sure you're referencinp
	// the pointer's value
	// (*p).firstName = newFirstName
}

// GOTCHA:
// a slice will not pass a copy
// and instead will immediately allow you to manipulate the
// original value
// because a slice is an abstraction on the array data structure,
// when passing a slice into a function, the slice value is copied
// but in the copied slice value, the pointer to the underlying
// array is also copied over exactly and therefore even the slice
// copy points to the original underlying array data structure
// a slice is a "reference type"
// more reference types that behave the same:
// maps, channels, pointers, functions

// an integer, float, string, bool, and struct
// will however pass a copy and therefore you need
// to reference the pointer in order to manipulate
// the original value
// these are "value types"
