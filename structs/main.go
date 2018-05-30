package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	//first way to create new person
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	//second way to define a new struct
	alex = person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
}
