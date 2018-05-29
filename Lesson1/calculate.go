package main

import (
	"fmt"
)

var numSlice []int

func newSlice(num int) []int {
	for i := 0; i <= num; i++ {
		numSlice = append(numSlice, i)
	}
	return numSlice
}

func printEvenOrOdd(numSlice []int) {
	for num := range numSlice {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
