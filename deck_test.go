package main

import "testing"

func TestNewDeck(t *testing.T){
	deck:=newDeck()
	if len(deck)!=16{
		t.Errorf("Expected length of 16 but get %v",len(deck))
	}
	if deck[0] !="Ace of Spades" {
		t.Errorf("Expected fist card is Ace of Spades but getting %v", deck[0])
	}
}