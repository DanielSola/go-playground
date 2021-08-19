package main

import "testing"
func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if(len(cards) != 16) {
		t.Errorf("Expected deck to be of length 16 but got %v", len(cards))
	}
}