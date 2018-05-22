//package=project=workspace
//package have two types (Executable and Reusable)
//Reusable : like code dependencies or library.(helper function or some code will used in future).
package main

import (
	"fmt"
)

func main() {
	cards := []string{"a", "b", "c"}
	for index, card := range cards {
		fmt.Println(card)
	}
}
