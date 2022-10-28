package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 168 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}