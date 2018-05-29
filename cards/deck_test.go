package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		//inject string formate
		t.Errorf("Excepted deck length of 16,but got %v", len(d))
	}
}
