package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	//go => when ever we are declaring multi line structures
	//every single line must have a comma,even if it is the last property declaration or the last
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.updateName("jimmy")
	jim.print()
}
