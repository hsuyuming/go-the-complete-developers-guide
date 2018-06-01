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
	fmt.Printf("%p\n", &p) //result =>0xc42001e140
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
	fmt.Printf("%p\n", &p) //result => 0xc42001e180
	p.print()              //result=> {firstName:jimmy lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
}

/*
go is what we refer to as a pass by value language
whenevet we pass some value into a function
go will take that value(date) and copy all of that data inside that struct
and then place it inside of new some new container(memory address)
*/
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
	fmt.Printf("%p\n", &jim) //result =>0xc42001e100
	jim.updateName("jimmy")
	//jim.print() //result=>{firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
}
