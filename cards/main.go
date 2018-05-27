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
	cards := newDeck()
	fmt.Println(cards.toString())
	/*
		greeting := "Hi there!"
		fmt.Println([]byte(greeting))
	*/
}
