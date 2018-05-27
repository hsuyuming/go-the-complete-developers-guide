/*
*****************************
Array v.s Slice (both must be defined with a data type)

Array(primitive data) => always going to be a very fixed length of records

Slice => like an array cna grow or shrink at will.

*****************************
*/

package main

func main() {
	//slice of type string
	cards := deck{"Ace og Diamonds", newCard()}
	//append return the new slice not modify original one
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
