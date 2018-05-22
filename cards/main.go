/*
*****************************
Array v.s Slice (both must be defined with a data type)

Array(primitive data) => always going to be a very fixed length of records

Slice => like an array cna grow or shrink at will.

*****************************
*/

package main

import (
	"fmt"
)

func main() {
	//slice of type string
	cards := []string{"Ace og Diamonds", newCard()}
	//append return the new slice not modify original one
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
