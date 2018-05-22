package main

import (
	"fmt"
)

//go type =>[bool,string,int,float64]

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

/*
func newCard() int {
	return 12
}
*/
