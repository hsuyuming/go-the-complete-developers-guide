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

/*
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
	fmt.Printf("%p\n", &p) //result => 0xc42001e180
	p.print()              //result=> {firstName:jimmy lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
}
*/

//*person => This is a type description,it means we're working with a pointer to a person
//a type of pointer that pointer a person
func (pointerToPerson *person) updateName(newFirstName string) {
	fmt.Printf("%p\n", &*pointerToPerson) //result ->0xc42001e100
	//star operator => give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
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
	//ambersand operator => give new the memory address of the value this variable is pointing at
	fmt.Printf("%p\n", &jim) //result =>0xc42001e100
	jim.updateName("jimmy")
	jim.print() //result=>{firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
}
