/*
*****************************
Array v.s Slice (both must be defined with a data type)

Array(primitive data) => always going to be a very fixed length of records

Slice => like an array cna grow or shrink at will.

*****************************
*/

package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
