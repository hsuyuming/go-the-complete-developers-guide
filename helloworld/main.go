//package=project=workspace
//package have two types (Executable and Reusable)
//Reusable : like code dependencies or library.(helper function or some code will used in future).
package main

//package apple

import "fmt"

func main() {
	fmt.Println(getTitle())
}

//go don't need to define function first(from quit)
func getTitle() string {
	return "Harry Potter"
}
